package request

import "fmt"

type CreateMessageRequest struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func (cm *CreateMessageRequest) Validate() error {
	if cm.ChatID == "" {
		return fmt.Errorf("chat_id cannot be empty")
	}
	return nil
}
