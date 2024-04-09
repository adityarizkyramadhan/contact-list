package request

import (
	"fmt"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	ContactQuery struct {
		Name  string
		Page  int
		Limit int
	}
	ContactCreate struct {
		FirstName    string              `json:"first_name"`
		LastName     string              `json:"last_name"`
		UserID       string              `json:"-"`
		PhoneNumbers []PhoneNumberCreate `json:"phone_numbers"`
	}

	ContactUpdate struct {
		ID                 string              `json:"-"`
		UserID             string              `json:"-"`
		FirstName          string              `json:"first_name"`
		LastName           string              `json:"last_name"`
		PhoneNumberUpdates []PhoneNumberUpdate `json:"phone_numbers"`
	}
)

func (c ContactCreate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.FirstName, validation.Required, validation.Length(0, 255)),
		validation.Field(&c.LastName, validation.Length(0, 255)),
		validation.Field(&c.PhoneNumbers, validation.Required),
	)
}

func (c ContactUpdate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.FirstName, validation.Length(0, 255)),
		validation.Field(&c.LastName, validation.Length(0, 255)),
	)
}

func (c ContactQuery) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Length(0, 255)),
		validation.Field(&c.Page, validation.Min(1)),
		validation.Field(&c.Limit, validation.Min(1)),
	)
}

func BuildContactQuery(name, page, limit string) (ContactQuery, error) {
	// Jika name kosong, maka akan diisi dengan string kosong
	var query ContactQuery
	var err error
	if name != "" {
		query.Name = name
	}
	if page != "" {
		query.Page, err = strconv.Atoi(page)
		if err != nil {
			query.Page = 1
		}
	}
	if limit != "" {
		query.Limit, err = strconv.Atoi(limit)
		if err != nil {
			query.Limit = 10
		}
	}
	if err := query.Validate(); err != nil {
		return query, fmt.Errorf("bad request: %v", err.Error())
	}
	return query, nil
}
