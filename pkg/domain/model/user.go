package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ID int64 `gorm:"primarykey,autoIncrement,column:id" json:"id"`

	Name string `json:"name"`
}
