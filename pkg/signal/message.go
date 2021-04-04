package signal

type BroadcastMessage struct {
	Uuid  string
	Room  string
	Event *Event
}

type SignalMessage struct {
	Uuid   string
	ToUuid string
	Room   string
	Event  *Event
}

func NewBroadcastMessage(client *Client, event *Event) *BroadcastMessage {
	return &BroadcastMessage{
		Uuid:  client.Uid,
		Room:  client.Room,
		Event: event,
	}
}

func NewSignalMessage(client *Client, toUuid string, event *Event) *SignalMessage {
	return &SignalMessage{
		Uuid:   client.Uid,
		Room:   client.Room,
		ToUuid: toUuid,
		Event:  event,
	}
}

func NewTypeEvent(client *Client, eventType string, name string, ) *Event {
	event := &Event{
		Type: eventType,
		Name: name,
		Data: make(map[string]interface{}),
	}
	event.Data["room"] = client.Room
	event.Data["id"] = client.Uid
	return event
}
