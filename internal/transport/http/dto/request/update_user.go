package request

import "fmt"

type UpdateUserInfo struct {
	Password *string `json:"new_password"`
	Login    *string `json:"new_login"`
}

func (uu *UpdateUserInfo) Validate() error {
	if uu.Login == nil && uu.Password == nil {
		return fmt.Errorf("no data for update")
	}
	if uu.Login != nil {
		if *uu.Login == "" {
			return fmt.Errorf("new login cannot be empty")
		}
	}
	if uu.Password != nil {
		if *uu.Password == "" {
			return fmt.Errorf("new password cannot be empty")
		}
	}

	return nil
}
