package models

import "database/sql"

type LogEntry struct {
	LogID       *int    `json:"log_id, omitempty"`
	TimeCreated *string `json:"time_created"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	OwnerID     *int    `json:"owner_id"`
	ProjectID   *int    `json:"project_id"`
}

type Task struct {
	TaskID      *int    `json:"task_id, omitempty"`
	TimeCreated *string `json:"time_created"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	OwnerID     *int    `json:"owner_id"`
	ProjectID   *int    `json:"project_id"`
	LogID       *int    `json:"log_id"`
	Status      *int    `json:"status"`
}

type Project struct {
	ProjectID   *int    `json:"project_id, omitempty"`
	TimeCreated *string `json:"time_created"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	OwnerID     *int    `json:"owner_id"`
	Status      *string `json:"status"`
}

type Owners struct {
	OwnerID    *int            `json:"owner_id, omitempty"`
	OwnerName  *string         `json:"owner_name"`
	OwnerEmail *string         `json:"owner_email"`
	LastLogin  *sql.NullString `json:"last_login_time"`
}
