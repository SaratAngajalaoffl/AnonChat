package handlers

import (
	"SaratAngajalaoffl/AnonChat/server/models"

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
			Err: errors.New("room not found"),
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

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	pName := ctx.Param("name")

	room, ok := runtime.Rooms[rId]

	if !ok {
		ctx.AbortWithError(http.StatusBadRequest, gin.Error{
			Err: errors.New("room not found"),
		})
		return
	}

	for _, p := range room.Participants {
		if p.Name == pName {
			ctx.AbortWithError(http.StatusBadRequest, gin.Error{
				Err: errors.New("username already taken"),
			})
			return
		}
	}

	c, err := runtime.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	defer c.Close()

	participant := models.Participant{
		Name: pName,
		Conn: c,
	}

	room.Participants = append(room.Participants, &participant)

	handleParticipantJoined(room, &participant)
	defer handleParticipantLeft(room, &participant)

	for {
		messageType, p, err := c.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		handleBroadcastMessage(room, &participant, string(p[:]), messageType)
	}
}
