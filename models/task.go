package models

import "time"

// Task represents the task model
type Task struct {
	ID          int       `json:"id"`
	UserId      int       `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
