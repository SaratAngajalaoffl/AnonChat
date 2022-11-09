package models

import "github.com/gorilla/websocket"

type Participant struct {
	Name string `json:"name"`
	Conn *websocket.Conn
}
