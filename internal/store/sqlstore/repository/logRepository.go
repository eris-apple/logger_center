package repository

import (
	"errors"
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/store"
	"github.com/eris-apple/logger_center/internal/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type LogRepository struct {
	DB *gorm.DB
}

func (ur *LogRepository) Create(l *models.Log) error {
	id := uuid.NewV4().String()

	log := models.Log{
		ID:        id,
		ChainID:   l.ChainID,
		ProjectID: l.ProjectID,
		Content:   l.Content,
		Timestamp: l.Timestamp,
		Level:     l.Level,
	}

	result := ur.DB.Table("logs").Create(&log).Scan(&l)
	if result.Error != nil {
		return store.ErrRecordNotCreated
	}

	return result.Error
}

func (ur *LogRepository) Search(projectID string, queryString string, filter *utils.Filter) ([]models.Log, error) {
	var logs []models.Log
	filter = utils.GetDefaultsFilter(filter)

	result := ur.DB.
		Table("logs").
		Find(&logs).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		Where("project_id = ? and content ILIKE ?", projectID, "%"+queryString+"%").
		Scan(&logs)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return nil, store.ErrRecordNotFound
	}

	return logs, result.Error
}

func (ur *LogRepository) FindAll(projectID string, filter *utils.Filter) ([]models.Log, error) {
	var logs []models.Log
	filter = utils.GetDefaultsFilter(filter)

	result := ur.DB.
		Table("logs").
		Find(&logs).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		Where("project_id = ?", projectID).
		Scan(&logs)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return nil, store.ErrRecordNotFound
	}

	return logs, result.Error
}

func (ur *LogRepository) FindById(id string) (*models.Log, error) {
	log := &models.Log{}

	result := ur.DB.Table("logs").Where("id = ?", id).First(log).Scan(&log)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return log, result.Error
}

func (ur *LogRepository) FindByChainId(chainID string, filter *utils.Filter) ([]models.Log, error) {
	var logs []models.Log

	filter = utils.GetDefaultsFilter(filter)

	result := ur.DB.
		Table("logs").
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		Where("chain_id = ?", chainID).
		Find(&logs).
		Scan(&logs)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return logs, result.Error
}

func (ur *LogRepository) Update(log *models.Log) error {
	result := ur.DB.Table("logs").Save(log).Scan(&log)
	if result.Error != nil {
		return store.ErrRecordNotUpdated
	}

	return result.Error
}

func (ur *LogRepository) Delete(log *models.Log) error {
	result := ur.DB.Table("logs").Delete(log)
	if result.Error != nil {
		return store.ErrRecordNotDeleted
	}

	return result.Error
}
