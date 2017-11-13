package models

type Task struct {
	TaskID      *int    `json:"task_id, omitempty"`
	Time        *string `json:"time"`
	Description *string `json:"description"`
	Status      *int    `json:"status"` // 0 for uncompleted, 1 for completed
}
