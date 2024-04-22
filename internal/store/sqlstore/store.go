package sqlstore

import (
	"github.com/aetherteam/logger_center/internal/store"
	"gorm.io/gorm"
)

type Store struct {
	DB                *gorm.DB
	userRepository    *UserRepository
	projectRepository *ProjectRepository
	sessionRepository *SessionRepository
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

	s.userRepository = &UserRepository{
		Store: s,
	}

	return s.userRepository
}

func (s *Store) Project() store.ProjectRepository {
	if s.projectRepository != nil {
		return s.projectRepository
	}

	s.projectRepository = &ProjectRepository{
		Store: s,
	}

	return s.projectRepository
}

func (s *Store) Session() store.SessionRepository {
	if s.sessionRepository != nil {
		return s.sessionRepository
	}

	s.sessionRepository = &SessionRepository{
		Store: s,
	}

	return s.sessionRepository
}
