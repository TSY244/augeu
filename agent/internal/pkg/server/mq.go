package server

type EventMessage struct {
	EventType string      `json:"event_type"`
	Data      interface{} `json:"data"`
}

type CmdMessage struct {
	Cmd string `json:"cmd"`
}
