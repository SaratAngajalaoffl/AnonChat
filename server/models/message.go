package models

type Message struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

type MessageData struct {
	Type string  `json:"type"`
	Data Message `json:"data"`
}
