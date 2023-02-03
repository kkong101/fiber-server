package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kkong101/fiber-server/config"
	db "github.com/kkong101/fiber-server/db/sqlc"
	"github.com/rs/zerolog/log"
	"time"
)

// main StartApplication
func main() {
	config, err := config.LoadConfig("./config/environment/")

	if err != nil {
		errors.New("CONFIG ERROR")
	}

	fmt.Println(config.Database.Driver)
	fmt.Println(config.Database.Source)

	conn, err := sql.Open(config.Database.Driver, config.Database.Source)
	if err != nil {
		fmt.Println(err)
	}
	queries := db.New(conn)

	if err := conn.Ping(); err != nil {
		log.Error().Msg("CONNECTION FAIL")
	}

	test := context.Background()

	test1 := db.CreateUserParams{
		Name:     "zzz",
		Phone:    "123",
		Birthday: time.Now(),
		Password: "123",
	}

	queries.CreateUser(test, test1)
	tt, err := queries.GetUser(test, "test")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("RESULT")
	fmt.Println(tt)
	fmt.Println(tt.Birthday)
	fmt.Println(tt.Password)

	fmt.Println()
}
