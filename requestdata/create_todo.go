package requestdata

type CreateTodo struct {
	Title           string `json:"title"`
	Priority        string `json:"priority"`
	ActivityGroupID int64  `json:"activity_group_id"`
}
