package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/MarcBeckRT/myapp-go/src/teamstar/model"
	"github.com/MarcBeckRT/myapp-go/src/teamstar/service"
)

func CreateTraining(w http.ResponseWriter, r *http.Request) {
	training, err := getTraining(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreateTraining(training); err != nil {
		log.Errorf("Error calling service CreateTraining: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, training)
}

func GetTrainings(w http.ResponseWriter, _ *http.Request) {
	trainings, err := service.GetTrainings()
	if err != nil {
		log.Errorf("Error calling service GetTrainings: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sendJson(w, trainings)
}

func GetTraining(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	training, err := service.GetTraining(id)
	if err != nil {
		log.Errorf("Failure retrieving training with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if training == nil {
		http.Error(w, "404 training not found", http.StatusNotFound)
		return
	}
	sendJson(w, training)
}

func DeleteTraining(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	training, err := service.DeleteTraining(id)
	if err != nil {
		log.Errorf("Failure deleting training with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if training == nil {
		http.Error(w, "404 training not found", http.StatusNotFound)
		return
	}
	sendJson(w, result{Success: "OK"})
}

func getTraining(r *http.Request) (*model.Training, error) {
	var training model.Training
	err := json.NewDecoder(r.Body).Decode(&training)
	if err != nil {
		log.Errorf("Can't serialize request body to training struct: %v", err)
		return nil, err
	}
	return &training, nil
}
