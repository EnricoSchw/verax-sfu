package signal

import (
	"github.com/pion/ion-sfu/pkg/sfu"
	"github.com/rs/xid"
	"log"
)

type Client struct {
	Uid            string
	Hub            *Hub
	Socket         *WebSocket
	Peer           *sfu.Peer
	Room           string
	isInConference bool
}

func NewClient(hub *Hub, ws *WebSocket, peer *sfu.Peer) *Client {
	client := &Client{
		Uid:            xid.New().String(),
		Hub:            hub,
		Socket:         ws,
		Peer:           peer,
		Room:           "",
		isInConference: false,
	}
	client.startListen()
	return client
}

func (client *Client) startListen() {
	client.Socket.On("room", func(event *Event) {
		if event.Name == "connect" {
			client.Room = event.Data["room"].(string)
			client.Hub.Register <- client
			client.send(NewTypeEvent(client, "room", "connectResponse"))
		}
	})

	client.Socket.On("conference", func(event *Event) {
		if event.Name == "join" {
			if client.Room != event.Data["room"].(string) {
				log.Printf("[Error] join confernce %v", event.Data)
			} else {
				peerJoinEvent := NewTypeEvent(client, "conference", "peerJoin")
				client.Hub.Broadcast <- NewBroadcastMessage(client, peerJoinEvent)
				client.send(NewTypeEvent(client, "conference", "joinResponse"))
			}
		}

		if event.Name == "leave" {
			if client.Room != event.Data["room"].(string) {
				log.Printf("[Error] join confernce %v", event.Data)
			} else {
				peerLeaveEvent := NewTypeEvent(client, "conference", "peerLeave")
				client.Hub.Broadcast <- NewBroadcastMessage(client, peerLeaveEvent)
				client.send(NewTypeEvent(client, "conference", "leaveResponse"))
			}
		}
	})

	client.Socket.On("signal", func(event *Event) {
		//if event.Data["sdp"] != nil {
		//	event.Data["sdp"] = event.Data["sdp"].(string)
		//}
		//if event.Data["candidate"] != nil {
		//	event.Data["candidate"] = event.Data["candidate"].(string)
		//}
		if client.Room != event.Data["room"].(string) {
			log.Printf("[Error] join confernce %v", event.Data)
		} else {
			signal:= NewTypeEvent(client, "signal", event.Name)
			signal.Signal = event.Signal
			client.Hub.Broadcast <- NewBroadcastMessage(client, signal)
		}
	})

	client.Socket.OnClose(func(event *Event) {
		disconnectEvent := NewTypeEvent(client, "system", "peerDisconnected")
		client.Hub.Broadcast <- NewBroadcastMessage(client, disconnectEvent)
		client.Hub.Unregister <- client
	})
}

func (client *Client) send(event *Event) {
	log.Printf("[MESSAGE] %v", event.Data)
	client.Socket.Out <- (event).Raw()
}
