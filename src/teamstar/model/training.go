package model

import "gorm.io/gorm"

type Training struct {
	gorm.Model
	Topic     string     `gorm:"notNull;size:40"`
	Date      string     `gorm:"notNull;size:40"`
	Feedbacks []Feedback `gorm:"foreignKey:TrainingID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
