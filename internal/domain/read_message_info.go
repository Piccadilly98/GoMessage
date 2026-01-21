package domain

import "time"

type ReadMessageDomain struct {
	ID          string
	ChatID      string
	SenderID    string
	RecipientID string
	CreatedDate time.Time
	IsReceived  bool
	Text        string
}
