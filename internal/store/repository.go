package store

import (
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/utils"
)

type UserRepository interface {
	Create(*models.User) error
	Search(*utils.Filter, string) (*[]models.User, error)
	FindAll(*utils.Filter, map[string]interface{}) (*[]models.User, error)
	FindById(string) (*models.User, error)
	FindByEmail(string) (*models.User, error)
	Update(*models.User) error
	Delete(*models.User) error
}

type SessionRepository interface {
	Create(*models.Session) error
	FindById(string) (*models.Session, error)
	FindByToken(string) (*models.Session, error)
	Delete(*models.Session) error
}

type ProjectRepository interface {
	Create(*models.Project) error
	FindAll(*utils.Filter) (*[]models.Project, error)
	FindById(string) (*models.Project, error)
	Update(*models.Project) error
	Delete(*models.Project) error
}

type LogRepository interface {
	Create(*models.Log) error
	FindAll(string, *utils.Filter) (*[]models.Log, error)
	FindById(string) (*models.Log, error)
	FindByProjectId(string, *utils.Filter) (*[]models.Log, error)
	FindByChainId(string, *utils.Filter) (*[]models.Log, error)
	Update(*models.Log) error
	Delete(*models.Log) error
}

type ServiceAccountRepository interface {
	Create(account *models.ServiceAccount) error
	FindAll(*utils.Filter) (*[]models.ServiceAccount, error)
	FindById(string) (*models.ServiceAccount, error)
	FindBySecret(string) (*models.ServiceAccount, error)
	FindByProjectId(string, *utils.Filter) (*[]models.ServiceAccount, error)
	Update(*models.ServiceAccount) error
	Delete(*models.ServiceAccount) error
}
