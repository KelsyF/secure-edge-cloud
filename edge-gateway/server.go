package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	addr string
	q    *MessageQueue
}

func NewServer(addr string, q *MessageQueue) *Server {
	return &Server{addr: addr, q: q}
}

func (s *Server) Start() {
	r := mux.NewRouter()
	r.HandleFunc("/event", s.handleEvent).Methods("POST")
	r.HandleFunc("/health", s.handleHealth).Methods("GET")

	log.Fatal(http.ListenAndServeTLS(s.addr, "cert.pem", "key.pem", r))
}
