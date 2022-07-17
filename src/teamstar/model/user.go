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

//type Users map[int]*User

//func (u Users) FindByName(name string) (*User, error) {
//	for _, user := range u {
//		if user.Name == name {
//			return user, nil
//		}
//	}
//	return nil, errors.New("USER_NOT_FOUND")
//}
