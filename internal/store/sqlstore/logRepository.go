package sqlstore

import (
	"errors"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/store"
	"github.com/aetherteam/logger_center/internal/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type LogRepository struct {
	Store *Store
}

func (ur *LogRepository) Create(l *models.Log) error {
	id := uuid.NewV4().String()

	log := models.Log{
		ID:        id,
		ChainID:   l.ChainID,
		ProjectID: l.ProjectID,
		Content:   l.Content,
		Level:     l.Level,
	}

	result := ur.Store.DB.Table("logs").Create(&log).Scan(&l)
	if result.Error != nil {
		return store.ErrRecordNotCreated
	}

	return result.Error
}

func (ur *LogRepository) FindAll(filter *utils.Filter) (*[]models.Log, error) {
	logs := &[]models.Log{}

	if filter.Limit == 0 {
		filter.Limit = 10
	}
	if filter.Order == "" {
		filter.Order = "id desc"
	}

	result := ur.Store.DB.
		Table("logs").
		Find(&logs).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		Scan(&logs)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return nil, store.ErrRecordNotFound
	}

	return logs, result.Error
}

func (ur *LogRepository) FindById(id string) (*models.Log, error) {
	log := &models.Log{}

	result := ur.Store.DB.Table("logs").Where("id = ?", id).First(log).Scan(&log)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return log, result.Error
}

func (ur *LogRepository) FindByChainId(chainID string) (*[]models.Log, error) {
	logs := &[]models.Log{}

	result := ur.Store.DB.Table("logs").Where("chain_id = ?", chainID).Find(logs).Scan(&logs)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return logs, result.Error
}

func (ur *LogRepository) Update(log *models.Log) error {
	result := ur.Store.DB.Table("logs").Save(log).Scan(&log)
	if result.Error != nil {
		return store.ErrRecordNotUpdated
	}

	return result.Error
}

func (ur *LogRepository) Delete(log *models.Log) error {
	result := ur.Store.DB.Table("logs").Delete(log)
	if result.Error != nil {
		return store.ErrRecordNotDeleted
	}

	return result.Error
}
