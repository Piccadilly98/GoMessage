package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Piccadilly98/GoMessage/internal/domain"
)

type UserPostrgres struct {
	db *sql.DB
}

func NewUserPostres(db *sql.DB) (*UserPostrgres, error) {
	if db == nil {
		return nil, fmt.Errorf("db cannot be nil\n")
	}

	return &UserPostrgres{
		db: db,
	}, nil
}

func (up *UserPostrgres) Create(ctx context.Context, req *domain.RegistrationUserDomain) (*domain.ReadUserDomain, error) {
	result := &domain.ReadUserDomain{}
	err := up.db.QueryRowContext(ctx, `
	INSERT INTO users(login, password_hash)
	VALUES($1, $2)
	RETURNING id, login, password_hash, created_date, updated_date;`,
		req.Login, req.PasswordHash).Scan(&result.ID, &result.Login, &result.PasswordHash, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return result, nil
}
