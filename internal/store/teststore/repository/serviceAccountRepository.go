package teststore_repository

import (
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/store"
	"github.com/eris-apple/logger_center/internal/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

type ServiceAccountRepository struct {
	ServiceAccounts map[string]*models.ServiceAccount
}

func (sar *ServiceAccountRepository) Search(projectID string, queryString string, filter *utils.Filter) ([]models.ServiceAccount, error) {
	filter = utils.GetDefaultsFilter(filter)

	var asa []models.ServiceAccount
	for _, sa := range sar.ServiceAccounts {
		if sa.ProjectID == projectID {
			asa = append(asa, *sa)
		}
	}

	if asa == nil {
		return nil, store.ErrRecordNotFound
	}

	fsa := utils.FilterArray(asa, filter)
	if fsa == nil {
		return nil, store.ErrRecordNotFound
	}

	var sa []models.ServiceAccount
	if queryString != "" {
		for _, l := range fsa {
			if strings.Contains(l.Name, queryString) {
				sa = append(sa, l)
			}
		}
	} else {
		sa = fsa
	}

	return sa, nil
}

func (sar *ServiceAccountRepository) FindAll(projectID string, filter *utils.Filter) ([]models.ServiceAccount, error) {
	filter = utils.GetDefaultsFilter(filter)

	var asa []models.ServiceAccount
	for _, sa := range sar.ServiceAccounts {
		if sa.ProjectID == projectID {
			asa = append(asa, *sa)
		}
	}

	if asa == nil {
		return nil, store.ErrRecordNotFound
	}

	sa := utils.FilterArray(asa, filter)
	if sa == nil {
		return nil, store.ErrRecordNotFound
	}

	return sa, nil
}

func (sar *ServiceAccountRepository) FindById(id string) (*models.ServiceAccount, error) {
	var sa *models.ServiceAccount
	for _, el := range sar.ServiceAccounts {
		if el.ID == id {
			sa = el
		}
	}

	if sa == nil {
		return nil, store.ErrRecordNotFound
	}

	return sa, nil
}

func (sar *ServiceAccountRepository) FindBySecret(secret string) (*models.ServiceAccount, error) {
	var sa *models.ServiceAccount
	for _, el := range sar.ServiceAccounts {
		if el.Secret == secret {
			sa = el
		}
	}

	if sa == nil {
		return nil, store.ErrRecordNotFound
	}

	return sa, nil
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

	sam := &models.ServiceAccount{
		ID:          id,
		ProjectID:   sa.ProjectID,
		Secret:      sa.Secret,
		Name:        sa.Name,
		Description: sa.Description,
	}

	sar.ServiceAccounts[sam.ID] = sam

	*sa = *sam

	return nil
}

func (sar *ServiceAccountRepository) Update(sa *models.ServiceAccount) error {
	csa := sar.ServiceAccounts[sa.ID]
	if csa == nil {
		return store.ErrRecordNotFound
	}

	sa.UpdatedAt = time.Now()
	sar.ServiceAccounts[sa.ID] = sa

	return nil
}

func (sar *ServiceAccountRepository) Delete(sa *models.ServiceAccount) error {
	csa := sar.ServiceAccounts[sa.ID]
	if csa == nil {
		return store.ErrRecordNotFound
	}

	sar.ServiceAccounts[sa.ID] = nil

	return nil
}
