package repository

import (
	"errors"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/store"
	"github.com/aetherteam/logger_center/internal/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type ServiceAccountRepository struct {
	DB *gorm.DB
}

func (sar *ServiceAccountRepository) FindAll(filter *utils.Filter) (*[]models.ServiceAccount, error) {
	sAccounts := &[]models.ServiceAccount{}
	filter = utils.GetDefaultsFilter(filter)

	result := sar.DB.
		Table("service_accounts").
		Find(&sAccounts).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
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

func (sar *ServiceAccountRepository) FindByProjectId(projectID string, filter *utils.Filter) (*[]models.ServiceAccount, error) {
	sAccount := &[]models.ServiceAccount{}
	filter = utils.GetDefaultsFilter(filter)

	result := sar.DB.
		Table("service_accounts").
		Where("project_id = ?", projectID).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		First(sAccount).
		Scan(&sAccount)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return sAccount, result.Error
}

func (sar *ServiceAccountRepository) Create(sA *models.ServiceAccount) error {
	id := uuid.NewV4().String()

	if validation.IsEmpty(sA.CreatedAt) {
		sA.CreatedAt = time.Time{}
	}

	if validation.IsEmpty(sA.UpdatedAt) {
		sA.UpdatedAt = time.Time{}
	}

	sAccount := models.ServiceAccount{
		ID:          id,
		ProjectID:   sA.ProjectID,
		Secret:      sA.Secret,
		Name:        sA.Name,
		Description: sA.Description,
	}

	result := sar.DB.Table("service_accounts").Create(&sAccount).Scan(&sA)
	if result.Error != nil {
		return store.ErrRecordNotCreated
	}

	return result.Error
}

func (sar *ServiceAccountRepository) Update(sA *models.ServiceAccount) error {
	result := sar.DB.Table("service_accounts").Save(sA).Scan(&sA)
	if result.Error != nil {
		return store.ErrRecordNotUpdated
	}

	return result.Error
}

func (sar *ServiceAccountRepository) Delete(sA *models.ServiceAccount) error {
	result := sar.DB.Table("service_accounts").Delete(sA)
	if result.Error != nil {
		return store.ErrRecordNotUpdated
	}

	return result.Error
}
