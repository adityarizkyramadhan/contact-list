package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PhoneNumber struct {
	ID        uuid.UUID      `json:"id" gorm:"type:varchar(36);primary_key;" db:"id"`
	Number    string         `json:"number" gorm:"type:varchar(255);not null;unique" db:"number"`
	ContactID uuid.UUID      `json:"contact_id" gorm:"type:varchar(36);not null;references:contacts;onUpdate:CASCADE;onDelete:CASCADE;" db:"contact_id"`
	UserID    uuid.UUID      `json:"user_id" gorm:"type:varchar(36);not null;references:users;onUpdate:CASCADE;onDelete:CASCADE;" db:"user_id"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" db:"deleted_at"`
}
