package main

import (
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

	//crate Admin
	service.CreateAdmin()

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

	log.Info("Starting My-Aktion API server")

	router := mux.NewRouter()
	router.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	router.HandleFunc("/logout", handler.LogoutHandler).Methods("POST")
	router.HandleFunc("/trainer/training", handler.CreateTraining).Methods("POST")
	router.HandleFunc("/trainer/users", handler.CreateUser).Methods("POST")
	router.HandleFunc("/player/trainings", handler.GetTrainings).Methods("GET")
	router.HandleFunc("/player/trainings/{id}", handler.GetTraining).Methods("GET")
	router.HandleFunc("/player/trainings/{id}", handler.DeleteTraining).Methods("DELETE")
	router.HandleFunc("/player/trainings/{id}/feedback", handler.AddFeedback).Methods("POST")
	//if err := http.ListenAndServe(":8080", router); err != nil {
	//	log.Fatal(err)
	//}
	log.Fatal(http.ListenAndServe(":8080", authorization.Authorizer(authEnforcer, model.Users{})(router)))

}
