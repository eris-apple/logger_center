package teststore

import (
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/store"
	teststore_repository "github.com/eris-apple/logger_center/internal/store/teststore/repository"
)

type Store struct {
	userRepository           *teststore_repository.UserRepository
	sessionRepository        *teststore_repository.SessionRepository
	logRepository            *teststore_repository.LogRepository
	projectRepository        *teststore_repository.ProjectRepository
	serviceAccountRepository *teststore_repository.ServiceAccountRepository
}

func New() *Store {
	return &Store{
		userRepository: &teststore_repository.UserRepository{
			Users: make(map[string]*models.User),
		},
		sessionRepository: &teststore_repository.SessionRepository{
			Sessions: make(map[string]*models.Session),
		},
		logRepository: &teststore_repository.LogRepository{
			Logs: make(map[string]*models.Log),
		},
		projectRepository: &teststore_repository.ProjectRepository{
			Projects: make(map[string]*models.Project),
		},
		serviceAccountRepository: &teststore_repository.ServiceAccountRepository{
			ServiceAccounts: make(map[string]*models.ServiceAccount),
		},
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &teststore_repository.UserRepository{
		Users: make(map[string]*models.User),
	}

	return s.userRepository
}

func (s *Store) Session() store.SessionRepository {
	if s.sessionRepository != nil {
		return s.sessionRepository
	}

	s.sessionRepository = &teststore_repository.SessionRepository{
		Sessions: make(map[string]*models.Session),
	}

	return s.sessionRepository
}

func (s *Store) Log() store.LogRepository {
	if s.logRepository != nil {
		return s.logRepository
	}

	s.logRepository = &teststore_repository.LogRepository{
		Logs: make(map[string]*models.Log),
	}

	return s.logRepository
}

func (s *Store) Project() store.ProjectRepository {
	if s.projectRepository != nil {
		return s.projectRepository
	}

	s.projectRepository = &teststore_repository.ProjectRepository{
		Projects: make(map[string]*models.Project),
	}

	return s.projectRepository
}

func (s *Store) ServiceAccount() store.ServiceAccountRepository {
	if s.serviceAccountRepository != nil {
		return s.serviceAccountRepository
	}

	s.serviceAccountRepository = &teststore_repository.ServiceAccountRepository{
		ServiceAccounts: make(map[string]*models.ServiceAccount),
	}

	return s.serviceAccountRepository
}
