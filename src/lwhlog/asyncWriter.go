package lwhlog

import (
	"QueueEvent"
	"lwhcontext"
	"os"
)

type AsyncWriter struct {
	HC   chan QueueEvent.HandleEvent
	File *os.File
}

func NewAsyncWriter() *AsyncWriter {
	writer := &AsyncWriter{HC: lwhcontext.HandleQueue.Queue}
	f, _ := os.OpenFile("./log4go.log", os.O_CREATE|os.O_RDWR, 0766)
	writer.File = f
	go writer.loopHandle()
	return writer
}

func (writer *AsyncWriter) loopHandle() {
	for {
		select {
		case event := <-writer.HC:
			msg := event.Msg
			//fmt.Print(msg)
			writer.File.WriteAt([]byte(msg), event.Start)
		}
	}
}
