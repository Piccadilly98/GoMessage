package auth

import (
	"fmt"
	"time"
)

const (
	MaxParallelGenerated = 6
)

type AuthService struct {
	secretKey            string
	durationJWT          time.Duration
	maxParallelGenerated int
	hashingCh            chan struct{}
}

func New(secretKey string, durationJWT time.Duration, maxParallelGenerated int) (*AuthService, error) {
	if secretKey == "" {
		return nil, fmt.Errorf("secretKey cannot be empty")
	}
	if durationJWT <= 0 {
		return nil, fmt.Errorf("durationJWT cannot be <= 0")
	}
	if maxParallelGenerated <= 0 || maxParallelGenerated > MaxParallelGenerated {
		maxParallelGenerated = MaxParallelGenerated
	}
	return &AuthService{
		secretKey:            secretKey,
		durationJWT:          durationJWT,
		maxParallelGenerated: maxParallelGenerated,
		hashingCh:            make(chan struct{}, maxParallelGenerated),
	}, nil
}
