package auth

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/argon2"
)

const (
	argonTime    uint32 = 5
	argonMemory  uint32 = 32 * 1024
	argonThreads uint8  = 1
	saltLength          = 16
	keyLength           = 32
)

func (as *AuthService) HashPassword(ctx context.Context, password string) (string, error) {

	select {
	case as.hashingCh <- struct{}{}:
	case <-ctx.Done():
		return "", fmt.Errorf("cpu load")
	}
	defer func() {
		tick := time.NewTicker(100 * time.Millisecond)
		select {
		case <-as.hashingCh:
			return
		case <-tick.C:
			return
		}
	}()
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	res := argon2.IDKey([]byte(password), salt, argonTime, argonMemory, argonThreads, keyLength)

	saltBase64 := base64.StdEncoding.EncodeToString(salt)
	hashBase64 := base64.StdEncoding.EncodeToString(res)
	return fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		argonMemory, argonTime, argonThreads, saltBase64, hashBase64), nil
}

func (ac *AuthService) CheckPassword(password, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 || parts[1] != "argon2id" || parts[2] != "v=19" {
		return false, fmt.Errorf("invalid argon2 hash format")
	}

	var time, memory, threads int
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		return false, err
	}
	salt, err := base64.StdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}
	expectedHash, err := base64.StdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}
	calculatedByte := argon2.IDKey([]byte(password), salt, uint32(time), uint32(memory), uint8(threads), uint32(len(expectedHash)))
	return bytes.Equal(calculatedByte, expectedHash), nil
}
