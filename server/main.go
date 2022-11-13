package main

import (
	"SaratAngajalaoffl/AnonChat/server/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/create-room", handlers.CreateRoom)
	r.GET("/get-room/:rId", handlers.GetRoom)
	r.GET("/join-room/:rId/:name", handlers.JoinRoom)

	r.Run(":8080")
}
