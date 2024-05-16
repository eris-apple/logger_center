package teststore_repository

import (
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/store"
	"github.com/eris-apple/logger_center/internal/utils"
	uuid "github.com/satori/go.uuid"
	"strings"
)

type LogRepository struct {
	Logs map[string]*models.Log
}

func (lr *LogRepository) Search(projectID string, queryString string, filter *utils.Filter) ([]models.Log, error) {
	filter = utils.GetDefaultsFilter(filter)

	var al []models.Log
	for _, l := range lr.Logs {
		if l.ProjectID == projectID {
			al = append(al, *l)
		}
	}

	if al == nil {
		return nil, store.ErrRecordNotFound
	}

	fLogs := utils.FilterArray(al, filter)
	if fLogs == nil {
		return nil, store.ErrRecordNotFound
	}

	var logs []models.Log
	if queryString != "" {
		for _, l := range fLogs {
			if strings.Contains(l.Content, queryString) {
				logs = append(logs, l)
			}
		}
	} else {
		logs = fLogs
	}

	return logs, nil
}

func (lr *LogRepository) FindAll(projectID string, filter *utils.Filter) ([]models.Log, error) {
	filter = utils.GetDefaultsFilter(filter)

	var al []models.Log
	for _, l := range lr.Logs {
		if l.ProjectID == projectID {
			al = append(al, *l)
		}
	}

	if al == nil {
		return nil, store.ErrRecordNotFound
	}

	logs := utils.FilterArray(al, filter)
	if logs == nil {
		return nil, store.ErrRecordNotFound
	}

	return logs, nil
}

func (lr *LogRepository) FindById(id string) (*models.Log, error) {
	var l *models.Log
	for _, el := range lr.Logs {
		if el.ID == id {
			l = el
		}
	}

	if l == nil {
		return nil, store.ErrRecordNotFound
	}

	return l, nil
}

func (lr *LogRepository) FindByChainId(chainID string, filter *utils.Filter) ([]models.Log, error) {
	var cl []models.Log
	for _, l := range lr.Logs {
		if l.ChainID == chainID {
			cl = append(cl, *l)
		}
	}

	if cl == nil {
		return nil, store.ErrRecordNotFound
	}

	logs := utils.FilterArray(cl, filter)
	if logs == nil {
		return nil, store.ErrRecordNotFound
	}

	return logs, nil
}

func (lr *LogRepository) Create(l *models.Log) error {
	id := uuid.NewV4().String()

	log := &models.Log{
		ID:        id,
		ChainID:   l.ChainID,
		ProjectID: l.ProjectID,
		Title:     l.Title,
		Error:     l.Error,
		Params:    l.Params,
		Content:   l.Content,
		Timestamp: l.Timestamp,
		Level:     l.Level,
	}

	lr.Logs[log.ID] = log

	*l = *log

	return nil
}

func (lr *LogRepository) Update(log *models.Log) error {
	l := lr.Logs[log.ID]
	if l == nil {
		return store.ErrRecordNotFound
	}

	lr.Logs[log.ID] = log

	return nil
}

func (lr *LogRepository) Delete(log *models.Log) error {
	l := lr.Logs[log.ID]
	if l == nil {
		return store.ErrRecordNotFound
	}

	lr.Logs[log.ID] = nil

	return nil
}
