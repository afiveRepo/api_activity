package entity

type Todo struct {
	BaseEntity
	Tittle          string `json:"tittle"`
	IsActive        int    `json:"is_active"`
	Priority        string `json:"priority"`
	ActivityGroupID int    `json:"activity_group_id"`
}
