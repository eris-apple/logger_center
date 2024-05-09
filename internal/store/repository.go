package store

import (
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/utils"
)

type UserRepository interface {
	Create(user *models.User) error
	Search(filter *utils.Filter, queryString string) (*[]models.User, error)
	FindAll(filter *utils.Filter, where map[string]interface{}) (*[]models.User, error)
	FindById(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(user *models.User) error
}

type SessionRepository interface {
	Create(session *models.Session) error
	FindById(id string) (*models.Session, error)
	FindByToken(token string) (*models.Session, error)
	Delete(session *models.Session) error
}

type ProjectRepository interface {
	Create(project *models.Project) error
	FindAll(filter *utils.Filter) (*[]models.Project, error)
	FindById(id string) (*models.Project, error)
	Update(project *models.Project) error
	Delete(project *models.Project) error
}

type LogRepository interface {
	Create(log *models.Log) error
	FindAll(queryString string, filter *utils.Filter) (*[]models.Log, error)
	FindById(id string) (*models.Log, error)
	FindByProjectId(projectID string, filter *utils.Filter) (*[]models.Log, error)
	FindByChainId(chainID string, filter *utils.Filter) (*[]models.Log, error)
	Update(log *models.Log) error
	Delete(log *models.Log) error
}

type ServiceAccountRepository interface {
	Create(serviceAccount *models.ServiceAccount) error
	FindAll(filter *utils.Filter) (*[]models.ServiceAccount, error)
	FindById(id string) (*models.ServiceAccount, error)
	FindBySecret(secret string) (*models.ServiceAccount, error)
	FindByProjectId(projectID string, filter *utils.Filter) (*[]models.ServiceAccount, error)
	Update(serviceAccount *models.ServiceAccount) error
	Delete(serviceAccount *models.ServiceAccount) error
}
