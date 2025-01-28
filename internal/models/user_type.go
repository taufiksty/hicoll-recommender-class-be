package models

import (
	"time"

	"gorm.io/gorm"
)

type UserType struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`

	Users []User `json:"users,omitempty" gorm:"foreignKey:UserTypeID"`
}

func (userType *UserType) TableName() string {
	return "user_types"
}

func (userType *UserType) BeforeUpdate(tx *gorm.DB) (err error) {
	userType.UpdatedAt = time.Now()
	return
}
