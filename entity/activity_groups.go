package entity

type ActivityGroups struct {
	BaseEntity
	Title string `json:"title"`
	Email string `json:"email"`
}
