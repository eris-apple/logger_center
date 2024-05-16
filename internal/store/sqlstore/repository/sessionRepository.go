package repository

import (
	"errors"
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/store"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SessionRepository struct {
	DB *gorm.DB
}

func (sr *SessionRepository) Create(s *models.Session) error {
	id := uuid.NewV4().String()
	session := models.Session{
		ID:       id,
		Token:    s.Token,
		UserID:   s.UserID,
		IsActive: s.IsActive,
	}

	result := sr.DB.Table("sessions").Create(&session).Scan(&s)

	return result.Error
}

func (sr *SessionRepository) FindById(id string) (*models.Session, error) {
	session := &models.Session{}

	result := sr.DB.Table("sessions").Where("id = ?", id).First(session).Scan(&session)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return session, result.Error
}

func (sr *SessionRepository) FindByToken(token string) (*models.Session, error) {
	session := &models.Session{}

	result := sr.DB.Table("sessions").Where("token = ?", token).First(session).Scan(&session)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return session, result.Error
}

func (sr *SessionRepository) Delete(session *models.Session) error {
	result := sr.DB.Table("sessions").Delete(session)
	if result.Error != nil {
		return store.ErrRecordNotFound
	}

	return result.Error
}
