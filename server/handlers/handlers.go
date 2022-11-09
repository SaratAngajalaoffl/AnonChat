package handlers

import (
	"SaratAngajalaoffl/AnonChat/server/models"

	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createRoom(ctx *gin.Context) {
	var newRoom models.Room

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
