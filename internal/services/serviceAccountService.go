package services

import (
	"github.com/aetherteam/logger_center/internal/config"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/store"
	"github.com/aetherteam/logger_center/internal/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type ServiceAccountService struct {
	ServiceAccountRepository store.ServiceAccountRepository
	ProjectService           ProjectService
}

func (sas *ServiceAccountService) FindAll(filter *utils.Filter) (*[]models.ServiceAccount, *config.APIError) {
	sAccounts, err := sas.ServiceAccountRepository.FindAll(filter)
	if err != nil {
		return nil, config.ErrServiceAccountsNotFound
	}

	return sAccounts, nil
}

func (sas *ServiceAccountService) FindById(id string) (*models.ServiceAccount, *config.APIError) {
	sAccount, err := sas.ServiceAccountRepository.FindById(id)
	if err != nil {
		return nil, config.ErrServiceAccountNotFound
	}

	return sAccount, nil
}

func (sas *ServiceAccountService) FindBySecret(secret string) (*models.ServiceAccount, *config.APIError) {
	sAccount, err := sas.ServiceAccountRepository.FindBySecret(secret)
	if err != nil {
		return nil, config.ErrServiceAccountNotFound
	}

	return sAccount, nil
}

func (sas *ServiceAccountService) FindByProjectId(projectID string, filter *utils.Filter) (*[]models.ServiceAccount, *config.APIError) {
	_, projectErr := sas.ProjectService.FindById(projectID)
	if projectErr != nil {
		return nil, config.ErrProjectNotFound
	}

	sAccounts, err := sas.ServiceAccountRepository.FindByProjectId(projectID, filter)
	if err != nil {
		return nil, config.ErrUserNotFound
	}

	return sAccounts, nil
}

func (sas *ServiceAccountService) Create(sAccount *models.ServiceAccount) (*models.ServiceAccount, *config.APIError) {
	_, projectErr := sas.ProjectService.FindById(sAccount.ProjectID)
	if projectErr != nil {
		return nil, config.ErrProjectNotFound
	}

	if err := sas.ServiceAccountRepository.Create(sAccount); err != nil {
		return nil, config.ErrInternalServerError
	}

	return sAccount, nil

}

func (sas *ServiceAccountService) Update(id string, usa *models.ServiceAccount) (*models.ServiceAccount, *config.APIError) {
	sAccount, err := sas.FindById(id)
	if err != nil {
		return nil, config.ErrServiceAccountNotFound
	}

	if validation.IsEmpty(usa.ProjectID) {
		usa.ProjectID = sAccount.ProjectID
	}

	if validation.IsEmpty(usa.Secret) {
		usa.Secret = sAccount.Secret
	}

	if validation.IsEmpty(usa.Name) {
		usa.Name = sAccount.Name
	}

	if validation.IsEmpty(usa.IsActive) {
		usa.IsActive = sAccount.IsActive
	}

	if validation.IsEmpty(usa.CreatedAt) {
		usa.CreatedAt = sAccount.CreatedAt
	}

	usa.UpdatedAt = time.Now()

	if err := sas.ServiceAccountRepository.Update(usa); err != nil {
		return nil, config.ErrInternalServerError
	}

	return usa, nil
}

func (sas *ServiceAccountService) Delete(id string) *config.APIError {
	sAccount, err := sas.FindById(id)
	if err != nil {
		return config.ErrServiceAccountNotFound
	}

	if err := sas.ServiceAccountRepository.Delete(sAccount); err != nil {
		return config.ErrInternalServerError
	}

	return nil
}

func NewServiceAccountService(usa store.ServiceAccountRepository, ps ProjectService) *ServiceAccountService {
	return &ServiceAccountService{
		ServiceAccountRepository: usa,
		ProjectService:           ps,
	}
}
