package entity

type Todo struct {
	BaseEntity
	Title           string `json:"title"`
	IsActive        int    `json:"is_active"`
	Priority        string `json:"priority"`
	ActivityGroupID int    `json:"activity_group_id"`
}
