package handlers

import (
	"SaratAngajalaoffl/AnonChat/server/models"
	"encoding/json"

	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var runtime models.App = models.App{
	Rooms: map[int]*models.Room{},
	Upgrader: websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	},
}

func CreateRoom(ctx *gin.Context) {
	var newRoom models.Room

	if err := ctx.BindJSON(&newRoom); err != nil {
		return
	}

	rId := runtime.GenerateRoomId()

	newRoom.Id = rId

	runtime.Rooms[rId] = &newRoom

	ctx.JSON(http.StatusOK, gin.H{
		"roomId": rId,
	})
}

func GetRoom(ctx *gin.Context) {

	rId, err := strconv.Atoi(ctx.Param("rId"))

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	room, ok := runtime.Rooms[rId]

	if !ok {
		ctx.AbortWithError(http.StatusBadRequest, gin.Error{
			Err: errors.New("Room not found"),
		})
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"room":         room,
		"population":   len(room.Participants),
		"participants": room.Participants,
	})
}

func JoinRoom(ctx *gin.Context) {

	w, r := ctx.Writer, ctx.Request

	rId, err := strconv.Atoi(ctx.Param("rId"))
	pName := ctx.Param("name")

	room, ok := runtime.Rooms[rId]

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

	c, err := runtime.Upgrader.Upgrade(w, r, nil)
	defer c.Close()

	participant := models.Participant{
		Name: pName,
		Conn: c,
	}

	room.Participants = append(room.Participants, &participant)
	defer room.RemoveParticipant(&participant)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	for _, p := range room.Participants {

		message := models.Message{
			Sender:  pName,
			Message: participant.Name + " has joined the chat",
		}

		encMsg, err := json.Marshal(message)

		if err != nil {
			log.Println(err)
			return
		}

		if err := p.Conn.WriteMessage(1, encMsg); err != nil {
			log.Println(err)
			return
		}
	}

	for {
		messageType, p, err := c.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		message := models.Message{
			Sender:  pName,
			Message: string(p[:]),
		}

		encMsg, err := json.Marshal(message)

		if err != nil {
			log.Println(err)
			return
		}

		for _, pt := range room.Participants {
			if err := pt.Conn.WriteMessage(messageType, encMsg); err != nil {
				log.Println(err)
				return
			}
		}
	}
}
