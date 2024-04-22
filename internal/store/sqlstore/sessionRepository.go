package sqlstore

import (
	"errors"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/store"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SessionRepository struct {
	Store *Store
}

func (pr *SessionRepository) Create(s *models.Session) error {
	id := uuid.NewV4().String()
	project := models.Session{
		ID:       id,
		Token:    s.Token,
		UserID:   s.UserID,
		IsActive: s.IsActive,
	}

	result := pr.Store.DB.Table("sessions").Create(&project).Scan(&s)

	return result.Error
}

func (pr *SessionRepository) FindById(id string) (*models.Session, error) {
	project := &models.Session{}

	result := pr.Store.DB.Table("sessions").Where("id = ?", id).First(project).Scan(&project)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return project, result.Error
}

func (pr *SessionRepository) FindByToken(token string) (*models.Session, error) {
	project := &models.Session{}

	result := pr.Store.DB.Table("sessions").Where("token = ?", token).First(project).Scan(&project)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return project, result.Error
}

func (pr *SessionRepository) Delete(session *models.Session) error {
	result := pr.Store.DB.Table("sessions").Delete(session)
	if result.Error != nil {
		return store.ErrRecordNotFound
	}

	return result.Error
}
