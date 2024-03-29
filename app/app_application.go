package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/kkong101/fiber-server/app/dto"
	err2 "github.com/kkong101/fiber-server/app/err"
	"github.com/kkong101/fiber-server/app/module/user/request"
	"github.com/kkong101/fiber-server/config"
	db "github.com/kkong101/fiber-server/db/sqlc"
	err1 "github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"reflect"
	"time"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

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

	app := fiber.New(fiber.Config{
		ErrorHandler: dto.ErrorHandler,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	app.Post("/test", func(ctx *fiber.Ctx) error {

		p := new(request.UserRequest)

		err := dto.ParseAndValidate(ctx, p)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(p.Phone)
		return nil
	})

	app.Get("/test2", func(ctx *fiber.Ctx) error {
		err = err2.NewCustomError("TEST123", err2.ValidationErrorSignatureInvalid)

		err1.Wrap(err, "???")

		fmt.Println("####")
		fmt.Println(reflect.TypeOf(err))
		if test, ok := err.(*err2.CustomError); ok {
			switch test.Errors {
			case err2.ValidationErrorSignatureInvalid:
				fmt.Println("SUCCESS")

			}
		}
		fmt.Println("####")

		return err
	})

	app.Listen("localhost:3001")

	fmt.Println("RESULT")
	fmt.Println(tt)
	fmt.Println(tt.Birthday)
	fmt.Println(tt.Password)
}
