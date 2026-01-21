package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Piccadilly98/GoMessage/internal/auth"
	_ "github.com/lib/pq"
)

func main() {
	// db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5433/messenger?sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pg, err := postgres.NewUserPostres(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// req := &domain.RegistrationUserDomain{
	// 	Login:        "user",
	// 	PasswordHash: "user",
	// }
	// res, err := pg.Create(context.Background(), req)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(*res)

	authSvc, err := auth.New("get_env", 100*time.Second, -1)
	if err != nil {
		log.Fatal(err.Error())
	}
	nowReal := time.Now()
	wg := sync.WaitGroup{}
	for i := range 500 {
		wg.Add(1)
		go func(i int) {
			_, err := authSvc.HashPassword(context.Background(), fmt.Sprintf("pass+%d%d", i, i))
			if err != nil {
				fmt.Printf("number - %d, err: %s\n", i, err.Error())
			}
			defer wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(time.Since(nowReal))
}
