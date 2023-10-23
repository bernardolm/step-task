package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ID        int64          `json:"id" gorm:"primarykey,autoIncrement"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`

	Name string `json:"name"`
}
