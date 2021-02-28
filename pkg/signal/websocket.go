package signal

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	WRITE_WAIT = 100 * time.Second

	// Maximum message size allowed from peer.
	MAX_MSG_SIZE = 500000000

	// Time allowed to read the next pong message from the peer.
	PONG_WAIT = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	PING_PERIOD = (PONG_WAIT * 9) / 10
)

var (
	newline = []byte{'\n'}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocket struct {
	Conn   *websocket.Conn
	Out    chan []byte
	In     chan []byte
	Events map[string]EventHandler
}

func NewWebSocket(w http.ResponseWriter, r *http.Request) (*WebSocket, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[ERROR | SOCKET CONNECT] %v", err)
		return nil, err
	}
	// conn.SetWriteDeadline(time.Now().Add(MSG_TIMEOUT))
	ws := &WebSocket{
		Conn:   conn,
		Out:    make(chan []byte, 256),
		In:     make(chan []byte),
		Events: make(map[string]EventHandler),
	}
	go ws.Reader()
	go ws.Writer()
	return ws, nil
}

func (ws *WebSocket) Reader() {
	defer func() {
		if action, ok := ws.Events["close"]; ok {
			action(&Event{
				Type: "system",
				Name: "close",
			})
		}
		log.Printf("[ClOSE] Reader Connection")
		ws.Conn.Close()
	}()
	ws.Conn.SetReadLimit(MAX_MSG_SIZE)
	ws.Conn.SetReadDeadline(time.Now().Add(PONG_WAIT))
	ws.Conn.SetPongHandler(func(string) error { ws.Conn.SetReadDeadline(time.Now().Add(PONG_WAIT)); return nil })
	for {
		_, message, err := ws.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("[ERROR] %v", err)
			}
			log.Printf("[ERROR] %v", err)
			break
		}
		event, err := NewEventFromRaw(message)
		if err != nil {
			log.Printf("[ERROR | MSG] %v", err)
		} else {
			log.Printf("[MSG] %v", event)
		}
		if action, ok := ws.Events[event.Type]; ok {
			action(event)
		}
	}
}

func (ws *WebSocket) Writer() {
	ticker := time.NewTicker(PING_PERIOD)
	defer func() {
		log.Printf("[ClOSE] Writer Connection")
		ticker.Stop()
	}()
	for {
		select {
		case message, ok := <-ws.Out:
			ws.Conn.SetWriteDeadline(time.Now().Add(WRITE_WAIT))
			if !ok {
				ws.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				log.Printf("Whats up? #{message)")
				return
			}
			w, err := ws.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			// Add queued chat messages to the current websocket message.
			n := len(ws.Out)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-ws.Out)
			}
			w.Close()

		case <-ticker.C:
			ws.Conn.SetWriteDeadline(time.Now().Add(WRITE_WAIT))
			if err := ws.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (ws *WebSocket) On(eventType string, action EventHandler) *WebSocket {
	ws.Events[eventType] = action
	return ws
}

func (ws *WebSocket) OnClose(action EventHandler) *WebSocket {
	ws.Events["close"] = action
	return ws
}
