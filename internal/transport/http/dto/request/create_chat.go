package request

import "fmt"

type CreateChatRequest struct {
	PartnerID string `json:"partrner_id"`
}

func (cc *CreateChatRequest) Validate() error {
	if cc.PartnerID == "" {
		return fmt.Errorf("partner id cannot be empty")
	}
	return nil
}
