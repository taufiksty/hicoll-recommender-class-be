package models

import (
	"time"

	"gorm.io/gorm"
)

type Class struct {
	ID                    uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name                  string         `json:"name" gorm:"type:varchar(255);not null"`
	Image                 *string        `json:"image,omitempty" gorm:"type:varchar(255)"`
	Thumbnail             *string        `json:"thumbnail,omitempty" gorm:"type:varchar(255)"`
	Description           *string        `json:"description,omitempty" gorm:"type:text"`
	MetaDescription       *string        `json:"meta_description,omitempty" gorm:"type:text"`
	Level                 string         `json:"level" gorm:"type:varchar(100);not null"`
	ClassCategoryID       uint           `json:"class_category_id" gorm:"type:not null"`
	ClassCategory         ClassCategory  `gorm:"foreignKey:ClassCategoryID;references:ID"`
	Tags                  *string        `json:"tags,omitempty" gorm:"type:varchar(255)"`
	Slug                  string         `json:"slug" gorm:"type:varchar(255);not null"`
	Method                *string        `json:"method,omitempty" gorm:"type:varchar(255)"`
	Media                 *string        `json:"media,omitempty" gorm:"type:varchar(255)"`
	PrefixCode            *string        `json:"prefix_code,omitempty" gorm:"type:varchar(255)"`
	Materials             *string        `json:"materials,omitempty" gorm:"type:varchar(255)"`
	CollaborationFeed     *string        `json:"collaboration_feed,omitempty" gorm:"type:varchar(255)"`
	InstructorID          *uint          `json:"instructor_id,omitempty" gorm:"type:int"`
	Instructor            User           `gorm:"foreignKey:InstructorID;references:ID"`
	LearningLink          *string        `json:"learning_link,omitempty" gorm:"type:varchar(255)"`
	ConsultancyLink       *string        `json:"consultancy_link,omitempty" gorm:"type:varchar(255)"`
	ConsultancySchedule   *string        `json:"consultancy_schedule,omitempty" gorm:"type:varchar(255)"`
	GroupChatLink         *string        `json:"group_chat_link,omitempty" gorm:"type:varchar(255)"`
	RegistrationCloseDate *string        `json:"registration_close_date,omitempty" gorm:"type:date"`
	Price                 string         `json:"price" gorm:"type:varchar(255);default:0"`
	IsDeleted             bool           `json:"is_deleted" gorm:"type:boolean;default:false"`
	CreatedAt             time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt             time.Time      `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt             gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	Users []User `gorm:"many2many:users_classes;" json:"users,omitempty"`
}

func (class *Class) TableName() string {
	return "classes"
}

func (class *Class) BeforeUpdate(tx *gorm.DB) (err error) {
	class.UpdatedAt = time.Now()
	return
}
