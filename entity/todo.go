package entity

type Todo struct {
	Title           string `json:"title"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
	ActivityGroupID int64  `json:"activity_group_id"`
	BaseEntity
}
