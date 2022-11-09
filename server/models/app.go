package models

import "github.com/gorilla/websocket"

type App struct {
	Rooms    map[int]*Room
	Upgrader websocket.Upgrader
}
