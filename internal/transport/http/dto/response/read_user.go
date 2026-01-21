package response

import "time"

type ReadUserResponse struct {
	ID          string     `json:"user_id"`
	Login       string     `json:"username"`
	CreatedDate time.Time  `json:"created_date"`
	UpdatedDate *time.Time `json:"updated_date,omitempty"`
}
