package lwhlog

import (
	qevent "QueueEvent"
	context "lwhcontext"
	"sync"
)

type LogConfig struct {
	FilePath   string
	QueueSize  int
	BufferSize int
}

type Logger struct {
	sync.RWMutex
	queueSize int
	fileName  string
	Queue     chan qevent.AccessEvent
	Append    *AsyncFileAppend
}

func NewLogConfig(file string, qsize int) *LogConfig {
	return &LogConfig{file, qsize, 1024}
}

func NewLog(c LogConfig) *Logger {
	aqueue := context.AccessQueue
	hqueue := context.HandleQueue

	fileAppend := NewAsyncFileAppend(aqueue.Queue, hqueue.Queue, c.FilePath)
	l := &Logger{
		queueSize: c.QueueSize,
		fileName:  c.FilePath,
		Queue:     aqueue.Queue,
		Append:    fileAppend,
	}
	return l
}

//var argss = []interface{}{1,"ssssssss",float64(3.2)}
//var ev = qevent.AccessEvent{}

func (l *Logger) Infof(format string, args ...interface{}) {
	ev := qevent.AccessEvent{Format: format, Args: args}
	l.Queue <- ev //qevent.AccessEvent{format, args}

	//ev.Format = format
	//ev.Args = args
	//l.Queue <- ev
}
