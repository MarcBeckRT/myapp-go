package model

type Role string

type User struct {
	ID   int
	Name string `gorm:"notNull;size:40"`
	Role string `gorm:"notNull;size:40"`
}
