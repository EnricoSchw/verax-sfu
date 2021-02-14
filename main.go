package main

import (
	"log"
	"net/http"
)

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Printf("Get New connection ...%w", w)
	ws, err := NewWebSocket(w, r)
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

}

//type longLatStruct struct {
//	Long float64 `json:"longitude"`
//	Lat  float64 `json:"latitude"`
//}
//
//var clients = make(map[*websocket.Conn]bool)
//var broadcast = make(chan *longLatStruct)

func main() {
	log.Printf("Start Server...")
	http.HandleFunc("/ws", wsEndpoint)
	log.Fatal(http.ListenAndServe(":8999", nil))
}
