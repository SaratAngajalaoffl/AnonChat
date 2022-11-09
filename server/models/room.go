package models

type Room struct {
	Id           int    `json:"id"`
	Topic        string `json:"topic"`
	Participants []*Participant
}
