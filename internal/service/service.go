package service

import (
	"github.com/Piccadilly98/GoMessage/internal/repository"
)

type Service struct {
	userRepo    repository.UserRepository
	chatRepo    repository.ChatRepository
	messageRepo repository.MessageRepository
	cache       repository.CacheRepository
}

func NewService(
	userRepo repository.UserRepository,
	chatRepo repository.ChatRepository,
	messageRepo repository.MessageRepository,
	cache repository.CacheRepository,
) *Service {
	return &Service{
		userRepo:    userRepo,
		chatRepo:    chatRepo,
		cache:       cache,
		messageRepo: messageRepo,
	}
}
