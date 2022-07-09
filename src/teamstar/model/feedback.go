package model

import "gorm.io/gorm"

type Status string

const (
	YES Status = "YES"
	NO  Status = "NO"
)

type Feedback struct {
	gorm.Model
	TrainingID uint
	Status     Status `gorm:"notNull;type:ENUM('YES','NO')"`
	Reason     string `gorm:"size:40"`
	User       User   `gorm:"embedded;embeddedPrefix:user_"`
}
