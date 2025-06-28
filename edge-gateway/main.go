package main

import (
	"log"
)

func main() {
	cfg := LoadConfig()

	q := NewMessageQueue()
	forwarder := NewForwarder(cfg.CloudURL, q)

	go forwarder.Start()

	server := NewServer(cfg.ListenAddr, q)
	log.Printf("Starting Edge Gateway on %s", cfg.ListenAddr)
	server.Start()
}
