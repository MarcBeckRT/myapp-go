package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/MarcBeckRT/myapp-go/src/teamstar/model"
	"github.com/MarcBeckRT/myapp-go/src/teamstar/service"
)

func AddFeedback(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	feedback, err := getFeedback(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: if the campaign doesn't exist, return 404 - don't show FK error
	err = service.AddFeedback(id, feedback)
	if err != nil {
		log.Errorf("Failure adding donation to campaign with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, feedback)
}

func getFeedback(r *http.Request) (*model.Feedback, error) {
	var feedback model.Feedback
	err := json.NewDecoder(r.Body).Decode(&feedback)
	if err != nil {
		log.Errorf("Can't serialize request body to feedback struct: %v", err)
		return nil, err
	}
	return &feedback, nil
}
