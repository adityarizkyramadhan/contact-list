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
