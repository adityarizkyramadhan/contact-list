package main

import (
	"fmt"

	"github.com/adityarizkyramadhan/contact-list/config/db"
	"github.com/adityarizkyramadhan/contact-list/contact"
	"github.com/adityarizkyramadhan/contact-list/middleware"
	"github.com/adityarizkyramadhan/contact-list/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load env: %v", err.Error()))
	}
	db, err := db.InitGorm()
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err.Error()))
	}
	log.Info().Msg("success connect to database")
	app := gin.New()
	log.Info().Msg("success create new gin app")
	app.Use(middleware.LogActivity())
	log.Info().Msg("success use log activity middleware")
	app.Use(middleware.Timeout())
	log.Info().Msg("success use timeout middleware")
	app.Use(middleware.Error())
	log.Info().Msg("success use error middleware")
	user := user.New(db, app)
	user.Init()
	log.Info().Msg("success init user")
	contact := contact.New(db, app)
	contact.Init()
	log.Info().Msg("success init contact")
	if err := app.Run(":8080"); err != nil {
		panic(fmt.Sprintf("failed to run server: %v", err.Error()))
	}
	log.Info().Msg("success run server")
}
