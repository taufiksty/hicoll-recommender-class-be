package models

import (
	"time"

	"gorm.io/gorm"
)

type UserClass struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id" gorm:"type:int;not null"`
	ClassID   uint      `json:"class_id" gorm:"type:int;not null"`
	Rating    uint      `json:"rating" gorm:"type:int"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`

	User  *User  `json:"users,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Class *Class `json:"classes,omitempty" gorm:"foreignKey:ClassID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (userClass *UserClass) TableName() string {
	return "users_classes"
}

func (userClass *UserClass) BeforeUpdate(tx *gorm.DB) (err error) {
	userClass.UpdatedAt = time.Now()
	return
}
