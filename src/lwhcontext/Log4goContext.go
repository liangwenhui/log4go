package lwhcontext

import queue "queue"

var AccessQueue *queue.AccessQueue = queue.NewAccessQueue(10000000)
var HandleQueue *queue.HandleQueue = queue.NewHandleEvent(10000000)
