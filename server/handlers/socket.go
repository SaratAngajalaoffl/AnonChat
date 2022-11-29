package handlers

import (
	"SaratAngajalaoffl/AnonChat/server/models"
	"encoding/json"
	"log"
)

func handleParticipantLeft(room *models.Room, participant *models.Participant) {
	room.RemoveParticipant(participant)

	for _, p := range room.Participants {

		message := models.Message{
			Sender:  "",
			Message: participant.Name + " has left the chat",
		}

		messageData := models.MessageData{
			Type: "MESSAGE",
			Data: message,
		}

		messageData2 := models.MessageData{
			Type: "PARTICIPANT_LEFT",
		}

		encMsg, err := json.Marshal(messageData)

		if err != nil {
			log.Println(err)
			return
		}

		if err := p.Conn.WriteMessage(1, encMsg); err != nil {
			log.Println(err)
			return
		}

		encMsg2, err := json.Marshal(messageData2)

		if err != nil {
			log.Println(err)
			return
		}

		if err := p.Conn.WriteMessage(1, encMsg2); err != nil {
			log.Println(err)
			return
		}
	}
}

func handleParticipantJoined(room *models.Room, participant *models.Participant) {
	for _, p := range room.Participants {

		message := models.Message{
			Sender:  "",
			Message: participant.Name + " has joined the chat",
		}

		messageData := models.MessageData{
			Type: "MESSAGE",
			Data: message,
		}

		messageData2 := models.MessageData{
			Type: "PARTICIPANT_JOINED",
		}

		encMsg, err := json.Marshal(messageData)

		if err != nil {
			log.Println(err)
			return
		}

		if err := p.Conn.WriteMessage(1, encMsg); err != nil {
			log.Println(err)
			return
		}

		encMsg2, err := json.Marshal(messageData2)

		if err != nil {
			log.Println(err)
			return
		}

		if err := p.Conn.WriteMessage(1, encMsg2); err != nil {
			log.Println(err)
			return
		}
	}
}

func handleBroadcastMessage(room *models.Room, participant *models.Participant, messageString string, messageType int) {
	message := models.Message{
		Sender:  participant.Name,
		Message: messageString,
	}

	messageData := models.MessageData{
		Type: "MESSAGE",
		Data: message,
	}

	encMsg, err := json.Marshal(messageData)

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
