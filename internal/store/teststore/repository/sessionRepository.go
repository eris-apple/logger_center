package teststore_repository

import (
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/store"
	uuid "github.com/satori/go.uuid"
)

type SessionRepository struct {
	Sessions map[string]*models.Session
}

func (sr *SessionRepository) FindById(id string) (*models.Session, error) {
	session := sr.Sessions[id]
	if session == nil {
		return nil, store.ErrRecordNotFound
	}

	return session, nil
}

func (sr *SessionRepository) FindByToken(token string) (*models.Session, error) {
	for _, session := range sr.Sessions {
		if session.Token == token {
			return session, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (sr *SessionRepository) Create(s *models.Session) error {
	id := uuid.NewV4().String()

	sr.Sessions[id] = &models.Session{
		ID:       id,
		Token:    s.Token,
		UserID:   s.UserID,
		IsActive: s.IsActive,
	}

	*s = *sr.Sessions[id]

	return nil
}

func (sr *SessionRepository) Delete(session *models.Session) error {
	checkSession := sr.Sessions[session.ID]
	if checkSession == nil {
		return store.ErrRecordNotFound
	}

	sr.Sessions[session.ID] = nil

	return nil
}
