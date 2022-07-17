package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/MarcBeckRT/myapp-go/src/teamstar/model"
	"github.com/MarcBeckRT/myapp-go/src/teamstar/service"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Can't serialize request body to campaign struct: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreateUser(&user); err != nil {
		log.Printf("Error calling service CreateUser: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	uid, err := getUId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := service.GetUser(uid)
	if err != nil {
		log.Errorf("Failure retrieving user with ID %v: %v", uid, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "404 user not found", http.StatusNotFound)
		return
	}
	sendJson(w, user)
}

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	users := service.GetUsers()
	//if err != nil {
	//	log.Errorf("Error calling service GetUsers: %v", err)
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	sendJson(w, users)
}
