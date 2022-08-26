package requestdata

import validation "github.com/go-ozzo/ozzo-validation/v4"

type UpdateActivityGroups struct {
	Title string `json:"title"`
}

func (c UpdateActivityGroups) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required),
	)
}
