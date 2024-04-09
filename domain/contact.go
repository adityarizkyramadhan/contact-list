package domain

import (
	"context"
	"time"

	"github.com/adityarizkyramadhan/contact-list/request"
	"github.com/adityarizkyramadhan/contact-list/response"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Contact struct {
		ID           uuid.UUID      `json:"id" gorm:"type:varchar(36);primary_key;" db:"id"`
		FirstName    string         `json:"first_name" gorm:"type:varchar(255);not null;" db:"first_name"`
		LastName     string         `json:"last_name" gorm:"type:varchar(255);" db:"last_name"`
		UserID       uuid.UUID      `json:"user_id" gorm:"type:varchar(36);no	t null;references:users;onUpdate:CASCADE;onDelete:CASCADE;" db:"user_id"`
		CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
		PhoneNumbers []*PhoneNumber `json:"phone_numbers" gorm:"foreignKey:ContactID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" db:"phone_numbers"`
		UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
		DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index" db:"deleted_at"`
	}
	ContactRepository interface {
		Delete(ctx context.Context, userID, id uuid.UUID) error
		Update(ctx context.Context, model *Contact) error
		FindAll(ctx context.Context, userID uuid.UUID, query request.ContactQuery) ([]Contact, int64, error)
		FindByID(ctx context.Context, userID, id uuid.UUID) (*Contact, error)
		CreateWithPhoneNumbers(ctx context.Context, model *Contact, phoneNumbers []*PhoneNumber) error
		Create(model *Contact) error
	}

	ContactUsecase interface {
		Create(ctx context.Context, model *request.ContactCreate) error
		FindByID(ctx context.Context, userID, id string) (*Contact, error)
		FindAll(ctx context.Context, userID string, query request.ContactQuery) (*response.FindAll, error)
		Update(ctx context.Context, model *request.ContactUpdate) error
		Delete(ctx context.Context, userID, id string) error
	}
)
