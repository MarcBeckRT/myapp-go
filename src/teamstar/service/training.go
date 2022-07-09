package service

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/MarcBeckRT/myapp-go/src/teamstar/db"
	"github.com/MarcBeckRT/myapp-go/src/teamstar/model"
)

func CreateTraining(training *model.Training) error {
	result := db.DB.Create(training)
	if result.Error != nil {
		return result.Error
	}
	log.Infof("Successfully stored new training with ID %v in database.", training.ID)
	log.Tracef("Stored: %v", training)
	return nil
}

func GetTraining(id uint) (*model.Training, error) {
	training := new(model.Training)
	result := db.DB.Preload("Feedbacks").First(training, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", training)
	return training, nil

}

func GetTrainings() ([]model.Training, error) {
	var trainings []model.Training
	result := db.DB.Preload("Feedbacks").Find(&trainings)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Tracef("Retrieved: %v", trainings)
	return trainings, nil
}

func DeleteTraining(id uint) (*model.Training, error) {
	training, err := GetTraining(id)
	if training == nil || err != nil {
		return training, err
	}
	result := db.DB.Delete(training)
	if result.Error != nil {
		return nil, result.Error
	}
	entry := log.WithField("ID", id)
	entry.Info("Successfully deleted training.")
	return training, nil
}
