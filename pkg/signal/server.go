package signal

import (
	"github.com/pion/ion-sfu/pkg/sfu"
	"log"
	"net/http"
)

type SignalServer struct {
	*sfu.Peer
}

func NewSignalServer() (*SignalServer, error) {
	server := &SignalServer{}
	return server, nil
}

func (server *SignalServer) Start() {
	//
	conf := sfu.Config{}
	s := sfu.NewSFU(conf)
	log.Printf("Start Server...")
	hub := NewHub()
	go hub.run()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		ws, err := NewWebSocket(writer, request)
		if err != nil {
			panic(err)
		}
		client:= NewClient(hub, ws, sfu.NewPeer(s))
		log.Printf("New Client connected: %v", client.Uid)
	})
	log.Printf("Server listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8999", nil))

}
