package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/MarcBeckRT/myapp-go/src/teamstar/authorization"
	"github.com/MarcBeckRT/myapp-go/src/teamstar/db"
	"github.com/MarcBeckRT/myapp-go/src/teamstar/handler"
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
	router.HandleFunc("/player/trainings", handler.GetTrainings).Methods("GET")
	router.HandleFunc("/player/trainings/{id}", handler.GetTraining).Methods("GET")
	router.HandleFunc("/player/trainings/{id}", handler.DeleteTraining).Methods("DELETE")
	router.HandleFunc("/player/trainings/{id}/feedback", handler.AddFeedback).Methods("POST")
	//if err := http.ListenAndServe(":8080", router); err != nil {
	//	log.Fatal(err)
	//}
	http.ListenAndServe(":8080", sessionManager.LoadAndSave(router))
	log.Fatal(http.ListenAndServe(":8080", authorization.Authorizer(authEnforcer, service.GetUsers())(router)))

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	idstring := r.FormValue("ID")
	id, err := strconv.Atoi(idstring)
	if err != nil { // ... handle error
		panic(err)
	}
	userID := id

	// First renew the session token...
	err2 := sessionManager.RenewToken(r.Context())
	if err2 != nil {
		http.Error(w, err2.Error(), 500)
		return
	}

	// Then make the privilege-level change.
	sessionManager.Put(r.Context(), "userID", userID)
}

//func loginHandler(users []model.User) http.HandlerFunc {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if err := sessionManager.RenewToken(r.Context()); err != nil {
//			log.Errorf("500 ERROR", w, err)
//			return
//		}
//		idstring := r.PostFormValue("id")
//		id, err := strconv.Atoi(idstring)
//		if err != nil {
//			// ... handle error
//			panic(err)
//		}
//		user, err := service.GetUser(id)
//		if err != nil {
//			log.Errorf("400 WRONG_CREDENTIALS", w, err)
//			return
//		}
//		// setup session
//
//		sessionManager.Put(r.Context(), "userID", user.ID)
//		sessionManager.Put(r.Context(), "role", user.Role)
//		log.Info("Successfull login", w)
//	})
//}

//func logoutHandler() http.HandlerFunc {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if err := sessionManager.RenewToken(r.Context()); err != nil {
//			log.Errorf("500 ERROR", w, err)
//			return
//		}
//		log.Info("Successfull logout", w)
//	})
//}

//func createAdmin() model.Users {
//	users := model.Users{}
//	users[1] = &model.User{
//		ID:   1,
//		Name: "admin",
//		Role: "trainer",
//	}
//	log.Info("created Admin with name=admin and Id=1")
//	return users
//}
