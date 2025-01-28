package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint      `json:"id" gorm:"primary_key;autoIncrement"`
	Fullname     string    `json:"fullname" gorm:"type:varchar(255);not null"`
	Email        string    `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	PhoneNumber  string    `json:"phone_number" gorm:"type:varchar(20);not null"`
	Gender       string    `json:"gender" gorm:"type:varchar(10);not null"`
	Birthdate    string    `json:"birthdate" gorm:"type:date;not null"`
	Description  *string   `json:"description,omitempty" gorm:"type:text"`
	LinkedinURL  *string   `json:"linkedin_url,omitempty" gorm:"type:varchar(255)"`
	Image        *string   `json:"image,omitempty" gorm:"type:varchar(255)"`
	Interests    *string   `json:"interests,omitempty" gorm:"type:varchar(255)"`
	IsActive     bool      `json:"is_active" gorm:"type:boolean;default:true"`
	IsFirstLogin bool      `json:"is_first_login" gorm:"type:boolean;default:false"`
	Password     string    `json:"password" gorm:"type:varchar(255);not null"`
	Token        *string   `json:"token,omitempty" gorm:"type:varchar(255)"`
	UserTypeID   uint      `json:"user_type_id" gorm:"type:int;not null"`
	UserType     UserType  `gorm:"foreignKey:UserTypeID;references:ID"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`

	Classes []Class `json:"classes,omitempty" gorm:"many2many:users_classes;"`
}

func (user *User) TableName() string {
	return "users"
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	user.UpdatedAt = time.Now()
	return
}
