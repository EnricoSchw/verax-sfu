package signal

import "encoding/json"

type EventHandler func(*Event)

type Event struct {
	Type   string                 `json:"type"`
	Name   string                 `json:"name"`
	Data   map[string]interface{} `json:"data"`
	Signal *json.RawMessage       `json:"signal,omitempty"`
}

type Signal struct {
	Sdp       string `json:"sdp"`
	Candidate string `json:"candidate"`
}

func NewEventFromRaw(rawData []byte) (*Event, error) {
	event := new(Event)
	err := json.Unmarshal(rawData, event)
	return event, err
}

func (e *Event) Raw() []byte {
	raw, _ := json.Marshal(e)
	return raw
}
