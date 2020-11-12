package queue

import "QueueEvent"

type AccessQueue struct {
	Queue chan QueueEvent.AccessEvent
	Size  int
}

func NewAccessQueue(size int) *AccessQueue {
	a := &AccessQueue{}
	q := make(chan QueueEvent.AccessEvent, size)
	a.Queue = q
	a.Size = size
	return a
}
