package response

import "time"

type ReadUserResponse struct {
	ID        string     `json:"user_id"`
	Login     string     `json:"username"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_date,omitempty"`
}
