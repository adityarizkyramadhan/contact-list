package domain

import (
	"context"
	"time"

	"github.com/adityarizkyramadhan/contact-list/request"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	User struct {
		ID          uuid.UUID      `json:"id" gorm:"type:varchar(36);primary_key;"`
		Username    string         `json:"username" gorm:"type:varchar(255);not null; unique"`
		Password    string         `json:"-" gorm:"type:varchar(255);not null"`
		CreatedAt   time.Time      `json:"createdAt" gorm:"autoCreateTime"`
		UpdatedAt   time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
		DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
		PhoneNumber []PhoneNumber  `json:"-" gorm:"foreignKey:user_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
		Contact     []Contact      `json:"-" gorm:"foreignKey:user_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	}
	UserRepository interface {
		Store(ctx context.Context, user *User) error
		FindByID(ctx context.Context, id uuid.UUID) (*User, error)
		FindByUsername(ctx context.Context, username string) (*User, error)
		Update(ctx context.Context, user *User) error
		Delete(ctx context.Context, id uuid.UUID) error
	}
	UserUsecase interface {
		Register(ctx context.Context, user *request.UserRegister) error
		Login(ctx context.Context, user *request.UserLogin) (string, error)
		FindByID(ctx context.Context, id string) (*User, error)
		FindByUsername(ctx context.Context, username string) (*User, error)
		Update(ctx context.Context, user *request.UserUpdate) error
		Delete(ctx context.Context, id string) error
	}
)

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
