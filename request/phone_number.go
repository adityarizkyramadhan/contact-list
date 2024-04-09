package request

import validation "github.com/go-ozzo/ozzo-validation"

type PhoneNumberCreate struct {
	Number string `json:"number"`
}

func (c PhoneNumberCreate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Number, validation.Required, validation.Length(0, 255)),
	)
}

type PhoneNumberOnlyCreate struct {
	ContactID string `json:"contact_id"`
	UserID    string `json:"user_id"`
	Number    string `json:"number"`
}

type PhoneNumberUpdate struct {
	ID     string `json:"id"`
	Number string `json:"number"`
}

func (c PhoneNumberUpdate) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ID, validation.Required),
		validation.Field(&c.Number, validation.Required, validation.Length(0, 255)),
	)
}
