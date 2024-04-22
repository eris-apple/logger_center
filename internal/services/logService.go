package services

import (
	"errors"
	"fmt"
	"github.com/aetherteam/logger_center/internal/config"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/store"
	"github.com/aetherteam/logger_center/internal/utils"
)

type LogService struct {
	LogRepository     store.LogRepository
	ProjectRepository store.ProjectRepository
}

func (ls LogService) FindAll(filter *utils.Filter) (*[]models.Log, error) {
	logs, err := ls.LogRepository.FindAll(filter)
	if err != nil {
		return nil, errors.New(config.ErrLogsNotFound)
	}

	return logs, nil
}

func (ls LogService) FindById(id string) (*models.Log, error) {
	log, err := ls.LogRepository.FindById(id)
	if err != nil {
		return nil, errors.New(config.ErrLogNotFound)
	}

	return log, nil
}

func (ls LogService) FindByChainId(id string) (*[]models.Log, error) {
	logs, err := ls.LogRepository.FindByChainId(id)
	if err != nil {
		return nil, errors.New(config.ErrLogNotFound)
	}

	return logs, nil
}

func (ls LogService) Create(log *models.Log) (*models.Log, error) {
	if err := ls.LogRepository.Create(log); err != nil {
		fmt.Print(err)
		return nil, errors.New(config.ErrInternalServerError)
	}

	return log, nil
}

func (ls LogService) Update(id string, updatedLog *models.Log) (*models.Log, error) {
	log, _ := ls.FindById(id)

	if updatedLog.Content == "" {
		updatedLog.Content = log.Content
	}

	if updatedLog.Level == "" {
		updatedLog.Level = log.Level
	}

	if updatedLog.ChainID == "" {
		updatedLog.ChainID = log.ChainID
	}

	if updatedLog.ProjectID == "" {
		updatedLog.ProjectID = log.ProjectID
	}

	if err := ls.LogRepository.Update(updatedLog); err != nil {
		return nil, errors.New(config.ErrInternalServerError)
	}

	return updatedLog, nil
}

func (ls LogService) Delete(id string) error {
	log, _ := ls.FindById(id)

	if err := ls.LogRepository.Delete(log); err != nil {
		return errors.New(config.ErrInternalServerError)
	}

	return nil
}

func NewLogService(lr store.LogRepository, pr store.ProjectRepository) LogService {
	return LogService{
		LogRepository:     lr,
		ProjectRepository: pr,
	}
}
