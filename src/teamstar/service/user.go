package service

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/MarcBeckRT/myapp-go/src/teamstar/model"
)

var (
	users     map[int]*model.User
	actUserid int = 2
)

func init() {
	users = make(map[int]*model.User)
}

func CreateAdmin() {
	users[1] = &model.User{
		ID:   1,
		Name: "admin",
		Role: "trainer",
	}
	log.Info("created Admin with name=admin and Id=1")
}

func CreateUser(user *model.User) error {
	user.ID = actUserid
	users[actUserid] = user
	actUserid += 1
	log.Printf("Successfully created new user with ID %v.", user.ID)
	return nil
}

func GetUserID(name string) (int, error) {
	var userID int
	user, err := FindByName(name)
	if err != nil {
		return 0, err //noch überarbeiten
		//log nachricht
	}
	userID = user.ID

	return userID, nil
}

func Exists(id int) bool {
	exists := false
	for _, user := range users {
		if user.ID == id {
			return true
		}
	}
	return exists
}

func FindByName(name string) (*model.User, error) {
	for _, user := range users {
		if user.Name == name {
			return user, nil
		}
	}
	return nil, errors.New("USER_NOT_FOUND")
}
