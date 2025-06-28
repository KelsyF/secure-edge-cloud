package main

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"secure-edge-cloud/edge-gateway/generated/messages"

	"google.golang.org/protobuf/proto"
)

type Forwarder struct {
	cloudURL string
	q        *MessageQueue
}

func NewForwarder(url string, q *MessageQueue) *Forwarder {
	return &Forwarder{cloudURL: url, q: q}
}

func (f *Forwarder) Start() {
	for {
		event, ok := f.q.Dequeue()
		if !ok {
			time.Sleep(500 * time.Millisecond)
			continue
		}

		if err := f.sendEvent(event); err != nil {
			log.Printf("Failed to send event, requeueing: %v", err)
			f.q.Enqueue(event)
			time.Sleep(1 * time.Second)
		}
	}
}

func (f *Forwarder) sendEvent(e Event) error {
	// Create Protobuf event from internal Event struct
	deviceEvent := &messages.DeviceEvent{
		DeviceID: e.DeviceID,
		Type:     e.Type,
		Payload:  e.Payload,
	}

	// Marshal to Protobuf bytes
	data, err := proto.Marshal(deviceEvent)
	if err != nil {
		return err
	}

	// Send as x-protobuf content type
	resp, err := http.Post(f.cloudURL, "applciation/x-protobuf", bytes.NewBuffer(data))
	if err != nil || resp.StatusCode >= 300 {
		return err
	}
	return nil
}
