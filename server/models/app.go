package models

import (
	"math/rand"

	"github.com/gorilla/websocket"
)

type App struct {
	Rooms    map[int]*Room
	Upgrader websocket.Upgrader
}

func (a *App) GenerateRoomId() int {
	var id int

	for {
		id = 100000 + rand.Int()%1000000

		if _, ok := a.Rooms[id]; !ok {
			break
		}
	}

	return id
}
