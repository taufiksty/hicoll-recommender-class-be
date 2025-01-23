package models

import (
	"time"

	"gorm.io/gorm"
)

type ClassCategory struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (classCategory *ClassCategory) TableName() string {
	return "class_categories"
}

func (classCategory *ClassCategory) BeforeUpdate(tx *gorm.DB) (err error) {
	classCategory.UpdatedAt = time.Now()
	return
}
