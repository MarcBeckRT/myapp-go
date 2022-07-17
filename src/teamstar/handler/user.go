package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/alexedwards/scs/v2"

	"github.com/MarcBeckRT/myapp-go/src/teamstar/model"
	"github.com/MarcBeckRT/myapp-go/src/teamstar/service"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var sessionManager *scs.SessionManager

	name := r.PostFormValue("name")

	// First renew the session token...
	err := sessionManager.RenewToken(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	userID, err := service.GetUserID(name)

	if err != nil {
		log.Errorf("Error calling service GetTrainings: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendJson(w, userID)

	// Then make the privilege-level change.
	sessionManager.Put(r.Context(), "userID", userID)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	var sessionManager *scs.SessionManager
	if err := sessionManager.RenewToken(r.Context()); err != nil {
		log.Errorf("Error", w, err)
		return
	}
	log.Info("Succes logout")

}

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
