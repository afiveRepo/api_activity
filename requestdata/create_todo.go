package requestdata

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateTodo struct {
	Title           string `json:"title"`
	Priority        string `json:"priority"`
	ActivityGroupID int64  `json:"activity_group_id"`
}

func (c CreateTodo) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required),
		validation.Field(&c.Priority, validation.Required, validation.In("Very-High", "High", "Medium", "Low", "Very-Low")),
		validation.Field(&c.ActivityGroupID, validation.Required),
	)
}
