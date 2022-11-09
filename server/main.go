package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Participant struct {
	Name string `json:"name"`
	conn *websocket.Conn
}

type Room struct {
	Id           int    `json:"id"`
	Topic        string `json:"topic"`
	participants []*Participant
}

var rooms map[int]*Room = map[int]*Room{}

var upgrader = websocket.Upgrader{} // use default option

func removeParticipant(room int, i *Participant) {
	s := rooms[room].participants

	for ind, p := range s {
		if p == i {
			s[ind] = s[len(s)-1]
			rooms[room].participants = s[:len(s)-1]
		}
	}
}

func generateRoomId() int {
	var id int

	for {
		id = 100000 + rand.Int()%1000000

		if _, ok := rooms[id]; !ok {
			break
		}
	}

	return id
}

func createRoom(ctx *gin.Context) {
	var newRoom Room

	if err := ctx.BindJSON(&newRoom); err != nil {
		return
	}

	rId := generateRoomId()

	newRoom.Id = rId

	rooms[rId] = &newRoom

	ctx.JSON(http.StatusOK, gin.H{
		"roomId": rId,
	})
}

func getRoom(ctx *gin.Context) {

	rId, err := strconv.Atoi(ctx.Param("rId"))

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	room, ok := rooms[rId]

	if !ok {
		ctx.AbortWithError(http.StatusBadRequest, gin.Error{
			Err: errors.New("Room not found"),
		})
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"room":         room,
		"population":   len(room.participants),
		"participants": room.participants,
	})
}

func joinRoom(ctx *gin.Context) {

	w, r := ctx.Writer, ctx.Request

	rId, err := strconv.Atoi(ctx.Param("rId"))
	pName := ctx.Param("name")

	room, ok := rooms[rId]

	if !ok {
		ctx.AbortWithError(http.StatusBadRequest, gin.Error{
			Err: errors.New("Room not found"),
		})
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c, err := upgrader.Upgrade(w, r, nil)
	defer c.Close()

	participant := Participant{
		Name: pName,
		conn: c,
	}

	room.participants = append(room.participants, &participant)
	defer removeParticipant(rId, &participant)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	for _, p := range room.participants {
		if p.conn != c {
			if err := p.conn.WriteMessage(1, []byte(participant.Name+" has joined the chat")); err != nil {
				log.Println(err)
				return
			}
		}
	}

	for {
		messageType, p, err := c.ReadMessage()

		fmt.Println(messageType, p)

		if err != nil {
			log.Println(err)
			return
		}

		for _, pt := range room.participants {
			if pt.conn != c {
				if err := pt.conn.WriteMessage(messageType, p); err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}

func main() {
	r := gin.Default()

	r.POST("/create-room", createRoom)
	r.GET("/get-room/:rId", getRoom)
	r.GET("/join-room/:rId/:name", joinRoom)

	r.Run(":8080")
}
