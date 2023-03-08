package models

import "time"

type Channel struct {
	Id          int64     `json:"id"`
	WorkspaceId int64     `json:"workspace_id"`
	Name        string    `json:"name"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}
