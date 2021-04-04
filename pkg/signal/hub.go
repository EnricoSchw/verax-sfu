package signal

import "log"

type Hub struct {
	rooms map[string]map[*Client]bool

	// Signaling messages from the clients.
	Broadcast chan *BroadcastMessage

	// Signaling messages from the clients.
	Signal chan *SignalMessage

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan *BroadcastMessage),
		Signal:     make(chan *SignalMessage),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		rooms:      make(map[string]map[*Client]bool),
	}
}

func (hub *Hub) run() {
	for {
		select {
		case client := <-hub.Register:
			log.Printf("[INFO] register client %v for room  %v", client.Uid, client.Room)
			clients := hub.rooms[client.Room]
			if clients == nil {
				clients = make(map[*Client]bool)
				hub.rooms[client.Room] = clients
			}
			hub.rooms[client.Room][client] = true

		case client := <-hub.Unregister:
			clientList := hub.rooms[client.Room]
			if clientList != nil {
				if _, ok := clientList[client]; ok {
					delete(clientList, client)
					close(client.Socket.Out)
					if len(clientList) == 0 {
						delete(hub.rooms, client.Room)
					}
				}
			}
		case message := <-hub.Broadcast:
			clientList := hub.rooms[message.Room]

			for client := range clientList {
				if client.Uid != message.Uuid {
					select {
					case client.Socket.Out <- message.Event.Raw():
					default:
						log.Printf("Hub close??: %v", client.Uid)
						close(client.Socket.Out)
						delete(clientList, client)
						if len(clientList) == 0 {
							delete(hub.rooms, message.Room)
						}
					}
				}
			}
		case signal := <-hub.Signal:
			clientList := hub.rooms[signal.Room]

			for client := range clientList {
				if client.Uid == signal.ToUuid {
					select {
					case client.Socket.Out <- signal.Event.Raw():
					default:
						log.Printf("Hub close??: %v", client.Uid)
						close(client.Socket.Out)
						delete(clientList, client)
						if len(clientList) == 0 {
							delete(hub.rooms, signal.Room)
						}
					}
				}
			}
		}
	}

}
