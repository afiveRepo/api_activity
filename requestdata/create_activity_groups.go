package requestdata

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateActicityGroups struct {
	Email string `json:"email"`
	Title string `json:"title"`
}

func (c CreateActicityGroups) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required),
		validation.Field(&c.Email, validation.Required, is.Email),
	)
}
