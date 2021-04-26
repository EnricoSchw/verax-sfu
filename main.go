package main

import (
	"github.com/EnricoSchw/verax-sfu/pkg/signal"
	"log"
)

func main() {
	server, err := signal.NewSignalServer()
	if err != nil {
		log.Printf("[ERROR | INIT Server] %v", err)
	} else {
		server.Start()
	}
}
