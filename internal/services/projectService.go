package services

import (
	"errors"
	"github.com/aetherteam/logger_center/internal/config"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/store"
	"github.com/aetherteam/logger_center/internal/utils"
	"gorm.io/gorm"
)

type ProjectService struct {
	ProjectRepository store.ProjectRepository
}

func (ps *ProjectService) FindAll(filter *utils.Filter) (*[]models.Project, *config.APIError) {
	projects, err := ps.ProjectRepository.FindAll(filter)
	if err != nil {
		return nil, config.ErrUsersNotFound
	}

	return projects, nil
}

func (ps *ProjectService) FindById(id string) (*models.Project, *config.APIError) {
	project, err := ps.ProjectRepository.FindById(id)
	if err != nil {
		return nil, config.ErrUsersNotFound
	}

	return project, nil
}

func (ps ProjectService) Create(project *models.Project) (*models.Project, *config.APIError) {
	if err := ps.ProjectRepository.Create(project); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, config.ErrUserAlreadyExist
		}

		return nil, config.ErrInternalServerError
	}

	return project, nil
}

func (ps ProjectService) Update(project *models.Project) (*models.Project, *config.APIError) {
	_, _ = ps.FindById(project.ID)

	if err := ps.ProjectRepository.Update(project); err != nil {
		return nil, config.ErrInternalServerError
	}

	return project, nil
}

func (ps ProjectService) Delete(project *models.Project) *config.APIError {
	_, _ = ps.FindById(project.ID)

	if err := ps.ProjectRepository.Delete(project); err != nil {
		return config.ErrInternalServerError
	}

	return nil
}

func NewProjectService(pr store.ProjectRepository) ProjectService {
	return ProjectService{
		ProjectRepository: pr,
	}
}
