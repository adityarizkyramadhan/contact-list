package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Contact struct {
	ID        uuid.UUID      `json:"id" gorm:"type:varchar(36);primary_key;" db:"id"`
	FirstName string         `json:"first_name" gorm:"type:varchar(255);not null;" db:"first_name"`
	LastName  string         `json:"last_name" gorm:"type:varchar(255);" db:"last_name"`
	UserID    uuid.UUID      `json:"user_id" gorm:"type:varchar(36);no	t null;references:users;onUpdate:CASCADE;onDelete:CASCADE;" db:"user_id"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" db:"deleted_at"`
}
