
package main

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Event struct {
	DeviceID string `json:"device_id"`
	Type     string `json:"type"`
	Payload  string `json: "payload"`
}

func (s *Server) handleEvent(w http.RepsonseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var event Event
	if err := json.Unmarshal(body. $event); err != nil {
		http.Error(w,m "Invalid JSON", http.StatusBadRequest)
		return
	}

	s.q.Enqueue(event)
	w.WriteHeader(http.StatusAccepted)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}