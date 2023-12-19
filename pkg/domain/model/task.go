package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ID int64 `gorm:"primarykey,autoIncrement,column:id" json:"id"`

	// StateID int64 `json:"state_id"`
	// State   State `gorm:"foreignKey:state_id,embedded"`

	// TaskID *int64 `json:"task_id"`
	// Task   *Task  `gorm:"foreignKey:task_id,embedded"`

	UserID int64
	User   User

	Description string `json:"description"`
}
