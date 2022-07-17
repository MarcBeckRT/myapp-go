package service

import (
	"fmt"

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
	log.Info("created Admin with name:admin and Id:1")
}

func CreateUser(user *model.User) error {
	user.ID = actUserid
	users[actUserid] = user
	actUserid += 1
	log.Printf("Successfully created new user with ID %v.", user.ID)
	return nil
}

func GetUser(uid int) (*model.User, error) {
	user := users[uid]
	if user == nil {
		return nil, fmt.Errorf("no user with ID %d", uid)
	}
	log.Tracef("Retrieved: %v", user)
	return user, nil
}

func GetUsers() []model.User {
	var userlist []model.User
	for _, user := range users {
		userlist = append(userlist, *user)
	}
	log.Tracef("Retrieved: %v", userlist)
	return userlist
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
