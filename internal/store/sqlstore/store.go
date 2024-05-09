package sqlstore

import (
	"github.com/aetherteam/logger_center/internal/store"
	"github.com/aetherteam/logger_center/internal/store/sqlstore/repository"
	"gorm.io/gorm"
)

type Store struct {
	DB                       *gorm.DB
	userRepository           *repository.UserRepository
	sessionRepository        *repository.SessionRepository
	projectRepository        *repository.ProjectRepository
	logRepository            *repository.LogRepository
	serviceAccountRepository *repository.ServiceAccountRepository
}

func New(db *gorm.DB) *Store {
	return &Store{
		DB: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &repository.UserRepository{
		DB: s.DB,
	}

	return s.userRepository
}

func (s *Store) Session() store.SessionRepository {
	if s.sessionRepository != nil {
		return s.sessionRepository
	}

	s.sessionRepository = &repository.SessionRepository{
		DB: s.DB,
	}

	return s.sessionRepository
}

func (s *Store) Project() store.ProjectRepository {
	if s.projectRepository != nil {
		return s.projectRepository
	}

	s.projectRepository = &repository.ProjectRepository{
		DB: s.DB,
	}

	return s.projectRepository
}

func (s *Store) Log() store.LogRepository {
	if s.logRepository != nil {
		return s.logRepository
	}

	s.logRepository = &repository.LogRepository{
		DB: s.DB,
	}

	return s.logRepository
}

func (s *Store) ServiceAccount() store.ServiceAccountRepository {
	if s.serviceAccountRepository != nil {
		return s.serviceAccountRepository
	}

	s.serviceAccountRepository = &repository.ServiceAccountRepository{
		DB: s.DB,
	}

	return s.serviceAccountRepository
}
