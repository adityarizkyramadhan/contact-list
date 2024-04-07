package repository

import (
	"context"

	"github.com/adityarizkyramadhan/contact-list/domain"
	"github.com/adityarizkyramadhan/contact-list/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type movie struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserRepository {
	return &movie{db}
}

func (m *movie) Store(ctx context.Context, user *domain.User) error {
	return m.db.WithContext(ctx).Create(user).Error
}

func (m *movie) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var user domain.User
	if err := m.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *movie) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	if err := m.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *movie) Update(ctx context.Context, user *domain.User) error {
	// cari user berdasarkan id
	var oldUser domain.User
	tx := utils.BeginTransaction(m.db).WithContext(ctx)
	defer utils.Rollback(tx, true)
	if err := tx.Where("id = ?", user.ID).First(&oldUser).Error; err != nil {
		return err
	}
	if user.Username != "" {
		oldUser.Username = user.Username
	}
	if user.Password != "" {
		oldUser.Password = user.Password
	}
	if err := tx.Save(&oldUser).Error; err != nil {
		return err
	}
	return tx.Commit().Error
}

func (m *movie) Delete(ctx context.Context, id uuid.UUID) error {
	tx := utils.BeginTransaction(m.db).WithContext(ctx)
	var rollback bool

	defer utils.Rollback(tx, rollback)

	if err := tx.Where("user_id = ?", id).Delete(&domain.PhoneNumber{}).Error; err != nil {
		return err
	}
	if err := tx.Where("user_id = ?", id).Delete(&domain.Contact{}).Error; err != nil {
		return err
	}
	if err := tx.Where("id = ?", id).Delete(&domain.User{}).Error; err != nil {
		return err
	}
	rollback = false
	return tx.Commit().Error
}
