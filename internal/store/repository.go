package store

import (
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/utils"
)

type UserRepository interface {
	Create(*models.User) error
	FindAll(*utils.Filter) (*[]models.User, error)
	FindById(string) (*models.User, error)
	FindByEmail(string) (*models.User, error)
	Update(*models.User) error
	Delete(*models.User) error
}

type ProjectRepository interface {
	Create(*models.Project) error
	FindAll(*utils.Filter) (*[]models.Project, error)
	FindById(string) (*models.Project, error)
	Update(*models.Project) error
	Delete(*models.Project) error
}

type SessionRepository interface {
	Create(*models.Session) error
	FindById(string) (*models.Session, error)
	FindByToken(string) (*models.Session, error)
	Delete(*models.Session) error
}
