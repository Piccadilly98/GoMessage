package domain

import "time"

type ReadUserDomain struct {
	ID           string
	Login        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}
