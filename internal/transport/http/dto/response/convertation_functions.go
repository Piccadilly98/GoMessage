package response

import (
	"fmt"

	"github.com/Piccadilly98/GoMessage/internal/domain"
)

func ToResponseDTO(entitie *domain.ReadUserDomain, info *TokenInfo, isAdnin bool) *ResultResponseWithToken {
	user := &ReadUserResponse{
		ID:        entitie.ID,
		Login:     entitie.Login,
		CreatedAt: entitie.CreatedAt,
	}
	if isAdnin {
		user.UpdatedAt = entitie.UpdatedAt
	}

	return &ResultResponseWithToken{
		User:  user,
		Token: info,
	}
}

func ToTokenInfoDTO(token, tokenType string, expiresIn int) (*TokenInfo, error) {
	if token == "" {
		return nil, fmt.Errorf("token cannot be empty")
	}
	if tokenType == "" {
		tokenType = BasicTokenType
	}

	return &TokenInfo{
		Token:     token,
		Type:      tokenType,
		ExpiresIn: expiresIn,
	}, nil
}

func ToReadChatResponse(entitie *domain.ReadChatDomain) *BasicReadChatResponse {
	return &BasicReadChatResponse{
		ID:        entitie.ID,
		UserID1:   entitie.UserID1,
		UserID2:   entitie.UserID2,
		CreatedAt: entitie.CreatedAt,
	}
}

func ToReadMessageResponse(entitie *domain.ReadMessageDomain, senderUsername string) *ReadMessageResponse {
	return &ReadMessageResponse{
		ID:             entitie.ID,
		ChatID:         entitie.ChatID,
		SenderUserName: senderUsername,
		SenderID:       entitie.SenderID,
		CreatedAt:      entitie.CreatedAt,
		Text:           entitie.Text,
	}
}
