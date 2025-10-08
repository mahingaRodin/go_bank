package models

import (
	"time"
)

type Task struct {
    ID          string    `json:"id"`
    Title       string    `json:"title" binding:"required"`
    Description string    `json:"description"`
    Status      string    `json:"status"` // "pending", "in-progress", "completed"
    Priority    string    `json:"priority"` // "low", "medium", "high"
    DueDate     time.Time `json:"due_date,omitempty"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type TaskUpdate struct {
    Title       *string    `json:"title,omitempty"`
    Description *string    `json:"description,omitempty"`
    Status      *string    `json:"status,omitempty"`
    Priority    *string    `json:"priority,omitempty"`
    DueDate     *time.Time `json:"due_date,omitempty"`
}