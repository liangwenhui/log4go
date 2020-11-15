package lwhlog

import (
	"QueueEvent"
	"bytes"
	"strconv"
	"strings"
	"sync"
	"time"
)

//var pLogFile, _ = os.OpenFile("./log4go.log", os.O_CREATE|os.O_RDWR, 0766)

type AsyncFileAppend struct {
	sync.RWMutex
	AC           chan QueueEvent.AccessEvent
	HC           chan QueueEvent.HandleEvent
	buffer       *bytes.Buffer
	buffersize   int
	index        int64
	lastSendTime time.Time
	maxWait      int64
	//B *bufio.Writer
	//file *os.File
}

func NewAsyncFileAppend(c chan QueueEvent.AccessEvent, c2 chan QueueEvent.HandleEvent, filePath string) *AsyncFileAppend {
	inst := &AsyncFileAppend{AC: c, HC: c2}
	inst.buffersize = 2048 * 4
	//pLogFile, _ := os.OpenFile("./log4go.log", os.O_CREATE|os.O_RDWR, 0766)
	//w := bufio.NewWriterSize(pLogFile, inst.buffersize)
	//inst.B = w
	inst.buffer = bytes.NewBuffer(make([]byte, 0))
	inst.index = 0
	inst.lastSendTime = time.Now()
	inst.maxWait = 200
	NewAsyncWriter(filePath)
	go inst.loopHandle()
	go inst.loopCheck()
	return inst
}

//消费loop AccessEvent
func (append *AsyncFileAppend) loopHandle() {
	for {
		select {
		case event := <-append.AC:
			msg := format(event.Format, event.Args)
			append.doAppend(msg)
			//fmt.Printf(event.Format,event.Args)
		}
	}
}

func (append *AsyncFileAppend) loopCheck() {
	ticker := time.Tick(200 * time.Millisecond)
	for {
		select {
		case <-ticker:
			last := append.lastSendTime
			milliseconds := time.Since(last).Milliseconds()
			if milliseconds >= append.maxWait {
				append.send()
			}
		}
	}

}

func (append *AsyncFileAppend) doAppend(msg string) {
	if len(msg) < 0 {
		return
	}
	buffer := append.buffer
	buffer.WriteString(msg)
	if buffer.Len() >= append.buffersize {
		//fmt.Print(buffer.String())
		str := buffer.String()
		buffer.Reset()
		append.HC <- QueueEvent.HandleEvent{Msg: str, Start: append.index}
		append.index += int64(len(str))
		append.lastSendTime = time.Now()
		//append.HC <- QueueEvent.HandleEvent{Msg:buffer.String()}
	}

}

func (append *AsyncFileAppend) send() {
	buffer := append.buffer
	append.lastSendTime = time.Now()
	if buffer.Len() > 0 {
		str := buffer.String()
		buffer.Reset()
		append.HC <- QueueEvent.HandleEvent{Msg: str, Start: append.index}
		append.index += int64(len(str))
	}
}

func format(format string, args []interface{}) string {
	if args == nil {
		return format
	}
	if len(args) == 0 {
		return format
	}
	//args := arg.([]interface{})
	for i := 0; i < len(args); i++ {
		switch args[i].(type) {
		case string:
			format = strings.Replace(format, "%s", args[i].(string), 1)
		case int:
			format = strings.Replace(format, "%d", strconv.Itoa(args[i].(int)), 1)
		case float32:
			format = strings.Replace(format, "%f", strconv.FormatFloat(float64(args[i].(float32)), 'f', 4, 32), 1)
		case float64:
			format = strings.Replace(format, "%f", strconv.FormatFloat(args[i].(float64), 'f', 4, 64), 1)
		}

	}

	return format
}
