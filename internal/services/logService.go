package services

import (
	"fmt"
	"github.com/eris-apple/logger_center/internal/config"
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/store"
	"github.com/eris-apple/logger_center/internal/utils"
)

type LogService struct {
	LogRepository     store.LogRepository
	ProjectRepository store.ProjectRepository
}

func (ls *LogService) Search(projectID string, queryString string, filter *utils.Filter) ([]models.Log, error) {
	logs, err := ls.LogRepository.Search(projectID, queryString, filter)
	if err != nil {
		return nil, config.ErrLogsNotFound
	}

	return logs, nil
}

func (ls *LogService) FindAll(projectID string, filter *utils.Filter) ([]models.Log, error) {
	logs, err := ls.LogRepository.FindAll(projectID, filter)
	if err != nil {
		return nil, config.ErrLogsNotFound
	}

	return logs, nil
}

func (ls *LogService) FindById(id string) (*models.Log, error) {
	log, err := ls.LogRepository.FindById(id)
	if err != nil {
		return nil, config.ErrLogNotFound
	}

	return log, nil
}

func (ls *LogService) FindByChainId(chainID string, filter *utils.Filter) ([]models.Log, error) {
	logs, err := ls.LogRepository.FindByChainId(chainID, filter)
	if err != nil {
		return nil, config.ErrLogNotFound
	}

	return logs, nil
}

func (ls *LogService) Create(log *models.Log) (*models.Log, error) {
	if _, err := ls.ProjectRepository.FindById(log.ProjectID); err != nil {
		return nil, config.ErrProjectNotFound
	}

	if err := ls.LogRepository.Create(log); err != nil {
		return nil, config.ErrInternalServerError
	}

	return log, nil
}

func (ls *LogService) Update(id string, updatedLog *models.Log) (*models.Log, error) {
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

	updatedLog.CreatedAt = log.CreatedAt

	if err := ls.LogRepository.Update(updatedLog); err != nil {
		fmt.Println(updatedLog)
		return nil, config.ErrInternalServerError
	}

	return updatedLog, nil
}

func (ls *LogService) Delete(id string) error {
	log, _ := ls.FindById(id)

	if err := ls.LogRepository.Delete(log); err != nil {
		return config.ErrInternalServerError
	}

	return nil
}

func NewLogService(lr store.LogRepository, pr store.ProjectRepository) LogService {
	return LogService{
		LogRepository:     lr,
		ProjectRepository: pr,
	}
}
