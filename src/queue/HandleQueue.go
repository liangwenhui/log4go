package queue

import "QueueEvent"

type HandleQueue struct {
	Queue chan QueueEvent.HandleEvent
	Size  int
}

func NewHandleEvent(size int) *HandleQueue {
	a := &HandleQueue{}
	q := make(chan QueueEvent.HandleEvent, size)
	a.Queue = q
	a.Size = size
	return a
}
