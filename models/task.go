package models

import "time"

type Task struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Title       string    `json:"title" binding:"required" bson:"title"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	Category    string    `json:"category" bson:"category"`
	Priority    string    `json:"priority" binding:"required" bson:"priority"`
	Deadline    time.Time `json:"deadline" bson:"deadline"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}