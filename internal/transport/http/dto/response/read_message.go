package response

import "time"

type ReadMessageResponse struct {
	ID             string    `json:"message_id"`
	ChatID         string    `json:"chat_id"`
	SenderID       string    `json:"sender_id"`
	SenderUserName string    `json:"sender_username"`
	Text           string    `json:"text"`
	CreatedAt      time.Time `json:"sending_at"`
	IsReceived     bool      `json:"is_received"`
}
