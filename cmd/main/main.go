package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Piccadilly98/GoMessage/internal/domain"
	"github.com/Piccadilly98/GoMessage/internal/repository/postgres"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5433/messenger?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	pg, err := postgres.NewUserPostres(db)
	if err != nil {
		log.Fatal(err)
	}
	req := &domain.RegistrationUserDomain{
		Login:        "user",
		PasswordHash: "user",
	}
	res, err := pg.Create(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*res)
}
