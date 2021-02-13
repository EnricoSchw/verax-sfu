package main

import (
	"log"
	"net/http"
)

func main()  {

	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		ws, err := NewWebSocket(writer, request)
		if err != nil {
			panic(err)
		}
		ws.On("message", func(event *Event) {
			log.Printf("[MESSAGE] %v", event.Data)
			ws.Out <- (&Event{
				Name: "response",
				Data: event.Data,
			}).Raw()
		})
	})

	http.ListenAndServe("8080", nil)
}
