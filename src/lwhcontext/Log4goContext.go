package lwhcontext

import queue "queue"

var AccessQueue *queue.AccessQueue = queue.NewAccessQueue(20000000)
var HandleQueue *queue.HandleQueue = queue.NewHandleEvent(10000000)
