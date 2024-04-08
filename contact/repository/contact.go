package repository

import (
	"context"
	"fmt"

	"github.com/adityarizkyramadhan/contact-list/domain"
	"github.com/adityarizkyramadhan/contact-list/request"
	"github.com/adityarizkyramadhan/contact-list/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type contact struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ContactRepository {
	return &contact{db}
}

func (c *contact) Create(model *domain.Contact) error {
	return c.db.Create(model).Error
}

func (c *contact) CreateWithPhoneNumbers(ctx context.Context, model *domain.Contact, phoneNumbers []*domain.PhoneNumber) error {
	tx := c.db.WithContext(ctx).Begin()
	rollback := true
	defer utils.Rollback(tx, rollback)
	if err := tx.Create(model).Error; err != nil {
		return fmt.Errorf("internal server error: %v", err.Error())
	}
	if len(phoneNumbers) > 0 {
		if err := tx.CreateInBatches(phoneNumbers, len(phoneNumbers)).Error; err != nil {
			return fmt.Errorf("internal server error: %v", err.Error())
		}
	}
	rollback = false
	return tx.Commit().Error
}

func (c *contact) FindByID(ctx context.Context, userID, id uuid.UUID) (*domain.Contact, error) {
	var contact domain.Contact
	if err := c.db.WithContext(ctx).Preload("PhoneNumbers").Where("user_id = ? AND id = ?", userID, id).First(&contact).Error; err != nil {
		return nil, fmt.Errorf("bad request: %v", err.Error())
	}
	return &contact, nil
}

func (c *contact) Update(ctx context.Context, model *domain.Contact) error {
	// cari dulu id contact yang akan diupdate
	tx := utils.BeginTransaction(c.db).WithContext(ctx)
	rollback := true
	defer utils.Rollback(tx, rollback)

	var contact domain.Contact
	if err := tx.Where("id = ?", model.ID).Where("user_id = ?", model.UserID).First(&contact).Error; err != nil {
		return fmt.Errorf("bad request: %v", err.Error())
	}

	// update contact dengan if else jika kosong jangan diupdate
	if model.FirstName != "" {
		contact.FirstName = model.FirstName
	}

	if model.LastName != "" {
		contact.LastName = model.LastName
	}

	// update contact
	if err := tx.Save(&contact).Error; err != nil {
		return fmt.Errorf("internal server error: %v", err.Error())
	}

	rollback = false
	return tx.Commit().Error
}

func (c *contact) Delete(ctx context.Context, userID, id uuid.UUID) error {
	// cari dulu apakah contact ada
	tx := utils.BeginTransaction(c.db).WithContext(ctx)
	rollback := true
	defer utils.Rollback(tx, rollback)
	// hitung adakah id contact yang akan dihapus
	var count int64
	if err := tx.Model(&domain.Contact{}).Where("id = ?", id).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return fmt.Errorf("internal server error: %v", err.Error())
	}
	if count == 0 {
		return fmt.Errorf("bad request: contact not found")
	}
	// hapus dulu phone number
	if err := tx.Where("contact_id = ?", id).Delete(&domain.PhoneNumber{}).Error; err != nil {
		return fmt.Errorf("internal server error: %v", err.Error())
	}
	// hapus contact
	if err := tx.Delete(&domain.Contact{}, id).Error; err != nil {
		return fmt.Errorf("internal server error: %v", err.Error())
	}
	rollback = false
	return tx.Commit().Error
}

func (c *contact) FindAll(ctx context.Context, userID uuid.UUID, query request.ContactQuery) ([]domain.Contact, error) {
	var contacts []domain.Contact
	tx := c.db.WithContext(ctx).Where("user_id = ?", userID)
	if query.Name != "" {
		tx = tx.Where("first_name LIKE ?", "%"+query.Name+"%")
	}
	if query.Page != 0 && query.Limit != 0 {
		tx = tx.Offset((query.Page - 1) * query.Limit).Limit(query.Limit)
	}
	if err := tx.Preload("PhoneNumbers").Find(&contacts).Error; err != nil {
		return nil, fmt.Errorf("bad request: %v", err.Error())
	}
	return contacts, nil
}