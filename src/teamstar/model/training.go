package model

import "gorm.io/gorm"

type Show string

//const (
//	OPEN Show = "OPEN"
//	HIDE Show = "HIDE"
//)

type Training struct {
	gorm.Model
	Topic   string `gorm:"notNull;size:40"`
	Content string `gorm:"size:100"`
	//ShowContent Show       `gorm:"notNull;type:ENUM('OPEN','HIDE')"`
	Date      string     `gorm:"notNull;size:10"`
	Feedbacks []Feedback `gorm:"foreignKey:TrainingID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User      User       `gorm:"embedded;embeddedPrefix:user_"`
}
