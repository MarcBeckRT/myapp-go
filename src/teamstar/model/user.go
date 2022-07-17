package model

type Role string

//const (
//	TRAINER Role = "TRAINER"
//	PLAYER  Role = "PLAYER"
//)

type User struct {
	ID   int
	Name string `gorm:"notNull;size:40"`
	//Role Role   `gorm:"notNull;type:ENUM('TRAINER','PLAYER')"`
	Role string `gorm:"notNull;size:40"`
}

type Users []User
