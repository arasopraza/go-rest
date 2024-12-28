package main

import (
	"fmt"
	"go-rest-api/configs/database"
	"go-rest-api/internal/validator"
	"go-rest-api/store/user"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	log "github.com/rs/zerolog/log"
)

func main() {
	e := echo.New()
	e.Validator = validator.NewValidator()

	err := godotenv.Load(".env")
	if err != nil {
		log.Err(err)
	}

	//db config
	database, err := database.ConnectDB()
	if err != nil {
		log.Err(err)
	}

	//create group
	group := e.Group("/api/v1")

	//user
	userRepo := user.NewRepository(database)
	userUsecase := user.NewUsecase(userRepo)
	user.UserHandler(group, userUsecase, userRepo)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
