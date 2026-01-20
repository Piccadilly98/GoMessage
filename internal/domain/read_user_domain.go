package domain

import "time"

type ReadUserDomain struct {
	ID           string
	Login        string
	PasswordHash string
	CreatedDate  time.Time
	UpdatedDate  *time.Time
}
