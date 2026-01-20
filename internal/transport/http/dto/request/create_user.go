package request

import "fmt"

type RegistrationUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ru *RegistrationUserRequest) Validate() error {
	if ru.Username == "" {
		return fmt.Errorf("invalid username: empty")
	}
	if ru.Password == "" {
		return fmt.Errorf("invalid password: empty")
	}
	return nil
}
