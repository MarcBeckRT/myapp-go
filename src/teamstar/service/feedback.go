package service

import (
	log "github.com/sirupsen/logrus"

	"github.com/MarcBeckRT/myapp-go/src/teamstar/db"
	"github.com/MarcBeckRT/myapp-go/src/teamstar/model"
)

func AddFeedback(trainingId uint, feedback *model.Feedback) error {
	feedback.TrainingID = trainingId
	result := db.DB.Create(feedback)
	if result.Error != nil {
		return result.Error
	}
	entry := log.WithField("ID", trainingId)
	entry.Info("Successfully added feedback to training.")
	entry.Tracef("Stored: %v", feedback)
	return nil
}
