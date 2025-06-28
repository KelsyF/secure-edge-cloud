package main

import (
	"io"
	"log"
	"net/http"

	"secure-edge-cloud/edge-gateway/generated/messages"

	"google.golang.org/protobuf/proto"
)

func (s *Server) handleEvent(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var event messages.DeviceEvent
	if err := proto.Unmarshal(body, &event); err != nil {
		http.Error(w, "Invalid Protobuf", http.StatusBadRequest)
		return
	}

	log.Printf("Received event from %s of type %s", event.DeviceId, event.Type)

	s.q.Enqueue(Event{
		DeviceID: event.DeviceId,
		Type:     event.Type,
		Payload:  event.Payload,
	})

	w.WriteHeader(http.StatusAccepted)
}
