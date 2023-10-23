package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	// gorm.Model

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ID        int64          `json:"id" gorm:"primarykey,autoIncrement"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`

	TaskID *int64 `json:"belongs_to"`
	Task   *Task  `json:"task" gorm:"foreignKey:TaskID"`

	UserID int64 `json:"user_id"`
	User   User  `json:"user" gorm:"foreignKey:UserID"`

	Description string `json:"description"`
	State       string `json:"state"`
}
