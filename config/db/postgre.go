package db

import (
	"os"

	"github.com/adityarizkyramadhan/contact-list/domain"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm() (*gorm.DB, error) {
	connection := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Error().Msgf("cant connect to database %s", err)
		return nil, err
	}
	db.AutoMigrate(new(domain.User), new(domain.Contact), new(domain.PhoneNumber))
	return db, nil
}
