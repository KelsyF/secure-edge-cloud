
package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Forwarder struct {
	cloudURL string
	q 		 *MessageQueue
}

func NewForwarder(url string, q *MessageQueue) *Forwarder {
	return &Forwarder{cloudURL: url, q: q}
}

func (f *Forwarder) Start() {
	for {
		event, ok := f.q.Dequeue()
		if !ok {
			time.Sleep( * time.Millisecond)
			continue
		}

		if err := f.sendEvent(event); err != nil {
			log.Pringf("Fialed to send event, requeueing: %v", err)
			f.q.Enqueue(event)
			time.Sleep(1 * time.Second)
		}
	}
}

func (f *Forwarder) sendEvent(e Event) error {
	data, _ := json.Marshal(e)
	resp, err := http.Post(f.cloudURL, "application/json", bytes.NewBuffer(data))
	if err != nill || resp.StatusCode >= 300 {
		return err
	}
	return nil
}