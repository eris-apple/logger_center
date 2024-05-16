package repository

import (
	"errors"
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/store"
	"github.com/eris-apple/logger_center/internal/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type ServiceAccountRepository struct {
	DB *gorm.DB
}

func (sar *ServiceAccountRepository) Search(projectID string, queryString string, filter *utils.Filter) ([]models.ServiceAccount, error) {
	var sAccounts []models.ServiceAccount
	filter = utils.GetDefaultsFilter(filter)

	result := sar.DB.
		Table("service_accounts").
		Find(&sAccounts).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		Where("project_id = ? and name ILIKE ?", projectID, "%"+queryString+"%").
		Scan(&sAccounts)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return nil, store.ErrRecordNotFound
	}

	return sAccounts, result.Error
}

func (sar *ServiceAccountRepository) FindAll(projectID string, filter *utils.Filter) ([]models.ServiceAccount, error) {
	var sAccounts []models.ServiceAccount
	filter = utils.GetDefaultsFilter(filter)

	result := sar.DB.
		Table("service_accounts").
		Find(&sAccounts).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		Where("project_id = ?", projectID).
		Scan(&sAccounts)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return nil, store.ErrRecordNotFound
	}

	return sAccounts, result.Error
}

func (sar *ServiceAccountRepository) FindById(id string) (*models.ServiceAccount, error) {
	sAccount := &models.ServiceAccount{}

	result := sar.DB.
		Table("service_accounts").
		Where("id = ?", id).
		First(sAccount).
		Scan(&sAccount)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return sAccount, result.Error
}

func (sar *ServiceAccountRepository) FindBySecret(secret string) (*models.ServiceAccount, error) {
	sAccount := &models.ServiceAccount{}

	result := sar.DB.
		Table("service_accounts").
		Where("secret = ?", secret).
		First(sAccount).
		Scan(&sAccount)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return sAccount, result.Error
}

func (sar *ServiceAccountRepository) Create(sa *models.ServiceAccount) error {
	id := uuid.NewV4().String()

	now := time.Now()

	if validation.IsEmpty(sa.CreatedAt) {
		sa.CreatedAt = now
	}

	if validation.IsEmpty(sa.UpdatedAt) {
		sa.UpdatedAt = now
	}

	sAccount := models.ServiceAccount{
		ID:          id,
		ProjectID:   sa.ProjectID,
		Secret:      sa.Secret,
		Name:        sa.Name,
		Description: sa.Description,
	}

	result := sar.DB.Table("service_accounts").Create(&sAccount).Scan(&sa)
	if result.Error != nil {
		return store.ErrRecordNotCreated
	}

	return result.Error
}

func (sar *ServiceAccountRepository) Update(sa *models.ServiceAccount) error {
	result := sar.DB.Table("service_accounts").Save(sa).Scan(&sa)
	if result.Error != nil {
		return store.ErrRecordNotUpdated
	}

	return result.Error
}

func (sar *ServiceAccountRepository) Delete(sa *models.ServiceAccount) error {
	result := sar.DB.Table("service_accounts").Delete(sa)
	if result.Error != nil {
		return store.ErrRecordNotUpdated
	}

	return result.Error
}
