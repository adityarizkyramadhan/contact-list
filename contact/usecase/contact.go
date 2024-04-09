package usecase

import (
	"context"
	"fmt"

	"github.com/adityarizkyramadhan/contact-list/domain"
	"github.com/adityarizkyramadhan/contact-list/request"
	"github.com/adityarizkyramadhan/contact-list/response"
	"github.com/google/uuid"
)

type contact struct {
	repoContact domain.ContactRepository
}

func New(repoContact domain.ContactRepository) domain.ContactUsecase {
	return &contact{repoContact}
}

func (c *contact) Create(ctx context.Context, model *request.ContactCreate) error {
	// masukkan data contact ke dalam domain.Contact
	userID, err := uuid.Parse(model.UserID)
	if err != nil {
		return fmt.Errorf("bad request: %v", err.Error())
	}
	contact := &domain.Contact{
		ID:        uuid.New(),
		FirstName: model.FirstName,
		LastName:  model.LastName,
		UserID:    userID,
	}
	// masukkan data phone number ke dalam domain.PhoneNumber
	var phoneNumbers []*domain.PhoneNumber
	for _, phoneNumber := range model.PhoneNumbers {
		phoneNumbers = append(phoneNumbers, &domain.PhoneNumber{
			ID:        uuid.New(),
			ContactID: contact.ID,
			Number:    phoneNumber.Number,
			UserID:    userID,
		})
	}
	// simpan data contact dan phone number ke dalam database
	if err := c.repoContact.CreateWithPhoneNumbers(ctx, contact, phoneNumbers); err != nil {
		return err
	}
	return nil
}

func (c *contact) FindByID(ctx context.Context, userID, id string) (*domain.Contact, error) {
	contactID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("bad request: %v", err.Error())
	}
	userIDParse, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("bad request: %v", err.Error())
	}
	contact, err := c.repoContact.FindByID(ctx, userIDParse, contactID)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func (c *contact) FindAll(ctx context.Context, userID string, query request.ContactQuery) (*response.FindAll, error) {
	userIDParse, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("bad request: %v", err.Error())
	}
	contacts, count, err := c.repoContact.FindAll(ctx, userIDParse, query)
	if err != nil {
		return nil, err
	}
	response := &response.FindAll{
		Data:       contacts,
		Pagination: response.NewPagination(query.Page, query.Limit, int(count)),
	}
	return response, nil
}

func (c *contact) Update(ctx context.Context, model *request.ContactUpdate) error {
	// masukkan data contact ke dalam domain.Contact
	if err := model.Validate(); err != nil {
		return fmt.Errorf("bad request: %v", err.Error())
	}
	userID, err := uuid.Parse(model.UserID)
	if err != nil {
		return fmt.Errorf("bad request: %v", err.Error())
	}
	contact := &domain.Contact{
		ID:        uuid.MustParse(model.ID),
		FirstName: model.FirstName,
		LastName:  model.LastName,
		UserID:    userID,
	}
	// simpan data contact ke dalam database
	if err := c.repoContact.Update(ctx, contact); err != nil {
		return err
	}
	return nil
}

func (c *contact) Delete(ctx context.Context, userID, id string) error {
	contactID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("bad request: %v", err.Error())
	}
	userIDParse, err := uuid.Parse(userID)
	if err != nil {
		return fmt.Errorf("bad request: %v", err.Error())
	}
	if err := c.repoContact.Delete(ctx, userIDParse, contactID); err != nil {
		return err
	}
	return nil
}
