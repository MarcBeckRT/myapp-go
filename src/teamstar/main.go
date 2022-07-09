package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/MarcBeckRT/myapp-go/src/teamstar/db"
	"github.com/MarcBeckRT/myapp-go/src/teamstar/handler"
	"github.com/gorilla/mux"
)

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
	log.Info("Starting My-Aktion API server")
	router := mux.NewRouter()
	router.HandleFunc("/training", handler.CreateTraining).Methods("POST")
	router.HandleFunc("/trainings", handler.GetTrainings).Methods("GET")
	router.HandleFunc("/trainings/{id}", handler.GetTraining).Methods("GET")
	router.HandleFunc("/trainings/{id}", handler.DeleteTraining).Methods("DELETE")
	router.HandleFunc("/trainings/{id}/feedback", handler.AddFeedback).Methods("POST")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
