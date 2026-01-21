package domain

import "time"

type ReadMessageDomain struct {
	ID          string
	ChatID      string
	SenderID    string
	RecipientID string
	CreatedAt   time.Time
	IsReceived  bool
	Text        string
}
