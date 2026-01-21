package response

import "time"

type BasicReadChatResponse struct {
	ID        string    `json:"chat_id"`
	UserID1   string    `json:"user_id_1"`
	UserID2   string    `json:"user_id_2"`
	CreatedAt time.Time `json:"created_at"`
}
