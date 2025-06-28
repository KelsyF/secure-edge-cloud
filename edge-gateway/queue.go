package main

import "sync"

type Event struct {
	DeviceID string
	Type     string
	Payload  string
}

type MessageQueue struct {
	mu    sync.Mutex
	queue []Event
}

func NewMessageQueue() *MessageQueue {
	return &MessageQueue{
		queue: make([]Event, 0),
	}
}

func (q *MessageQueue) Enqueue(e Event) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, e)
}

func (q *MessageQueue) Dequeue() (Event, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.queue) == 0 {
		return Event{}, false
	}
	e := q.queue[0]
	q.queue = q.queue[1:]
	return e, true
}
