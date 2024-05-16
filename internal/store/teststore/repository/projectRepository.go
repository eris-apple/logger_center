package teststore_repository

import (
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/store"
	"github.com/eris-apple/logger_center/internal/utils"
	uuid "github.com/satori/go.uuid"
)

type ProjectRepository struct {
	Projects map[string]*models.Project
}

func (pr *ProjectRepository) Create(p *models.Project) error {
	id := uuid.NewV4().String()
	*p = models.Project{
		ID:          id,
		Name:        p.Name,
		Prefix:      p.Prefix,
		IsActive:    p.IsActive,
		Description: p.Description,
	}

	pr.Projects[id] = p

	return nil
}

func (pr *ProjectRepository) FindAll(filter *utils.Filter, where map[string]interface{}) ([]models.Project, error) {
	filter = utils.GetDefaultsFilter(filter)

	var ap []models.Project
	for _, p := range pr.Projects {
		ap = append(ap, *p)
	}

	fProjects := utils.FilterArray(ap, filter)
	if fProjects == nil {
		return nil, store.ErrRecordNotFound
	}

	var projects []models.Project
	if len(where) != 0 {
		for _, p := range fProjects {
			if where["Name"] == p.Name {
				projects = append(projects, p)
			}
			if where["prefix"] == p.Prefix {
				projects = append(projects, p)
			}
			if where["is_active"] == p.IsActive {
				projects = append(projects, p)
			}
			if where["created_at"] == p.CreatedAt {
				projects = append(projects, p)
			}
		}
	} else {
		projects = fProjects
	}

	return projects, nil
}

func (pr *ProjectRepository) FindById(id string) (*models.Project, error) {
	project := pr.Projects[id]
	if project == nil {
		return nil, store.ErrRecordNotFound
	}

	return project, nil
}

func (pr *ProjectRepository) Update(project *models.Project) error {
	p := pr.Projects[project.ID]
	if p == nil {
		return store.ErrRecordNotFound
	}

	pr.Projects[project.ID] = project

	return nil
}

func (pr *ProjectRepository) Delete(project *models.Project) error {
	p := pr.Projects[project.ID]
	if p == nil {
		return store.ErrRecordNotFound
	}

	pr.Projects[project.ID] = nil

	return nil
}
