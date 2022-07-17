package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/MarcBeckRT/myapp-go/src/teamstar/authorization"
	"github.com/MarcBeckRT/myapp-go/src/teamstar/db"
	"github.com/MarcBeckRT/myapp-go/src/teamstar/handler"
	"github.com/MarcBeckRT/myapp-go/src/teamstar/model"
	"github.com/MarcBeckRT/myapp-go/src/teamstar/service"

	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
	"github.com/casbin/casbin"
	"github.com/gorilla/mux"
)

var sessionManager *scs.SessionManager

func init() {
	// init logger
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Info("Log level not specified, set default to: INFO")
		log.SetLevel(log.InfoLevel)
	}
	log.SetLevel(level)

	//init database
	db.Init()

}

func main() {
	// setup casbin auth rules
	authEnforcer, err := casbin.NewEnforcerSafe("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	// setup session store
	sessionManager = scs.New()
	sessionManager.IdleTimeout = 30 * time.Minute
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = true
	sessionManager.Store = memstore.New()

	//crate Admin
	service.CreateAdmin()

	log.Info("Starting My-Aktion API server")

	router := mux.NewRouter()
	//router.HandleFunc("/login", loginHandler(users))
	//router.HandleFunc("/logout", logoutHandler())
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/userlist", handler.GetUsers).Methods("GET")
	router.HandleFunc("/trainer/training", handler.CreateTraining).Methods("POST")
	router.HandleFunc("/trainer/users", handler.CreateUser).Methods("POST")
	router.HandleFunc("/trainer/trainings/{id}", handler.DeleteTraining).Methods("DELETE")
	router.HandleFunc("/player/trainings", handler.GetTrainings).Methods("GET")
	router.HandleFunc("/player/trainings/{id}", handler.GetTraining).Methods("GET")
	router.HandleFunc("/player/trainings/{id}/feedback", handler.AddFeedback).Methods("POST")

	http.ListenAndServe(":8080", sessionManager.LoadAndSave(router))
	userField := &model.User{}
	http.ListenAndServe(":8080", authorization.Authorizer(authEnforcer, userField)(router))

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

type myInfo struct {
	ID int `json:"id"`
	//Name string `json:"name"`
	//Role string `json:"role"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	//get user Id from http.Request
	var info myInfo
	r.ParseForm()
	//idstring := r.FormValue("ID")
	data, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	json.Unmarshal(data, &info)
	//fmt.Println(string(data))
	fmt.Println(info) // just for checking

	userID := info.ID

	// First renew the session token
	err := sessionManager.RenewToken(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Then make the privilege-level change.
	sessionManager.Put(r.Context(), "userID", userID)
	log.Infof("Successfully loged in user with Id: %d", userID)
}

//func logoutHandler() http.HandlerFunc {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if err := sessionManager.RenewToken(r.Context()); err != nil {
//			log.Errorf("500 ERROR", w, err)
//			return
//		}
//		log.Info("Successfull logout", w)
//	})
//}
