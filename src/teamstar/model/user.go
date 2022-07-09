package model

type Role string

const (
	TRAINER Role = "TRAINER"
	PLAYER  Role = "PLAYER"
)

type User struct {
	UserID uint
	Name   string `gorm:"notNull;size:40"`
	Role   Role   `gorm:"notNull;type:ENUM('TRAINER','PLAYER')"`
}
