package sqlstore

import (
	"errors"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/store"
	"github.com/aetherteam/logger_center/internal/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	Store *Store
}

func (pr *ProjectRepository) Create(p *models.Project) error {
	id := uuid.NewV4().String()
	project := models.Project{
		ID:       id,
		Name:     p.Name,
		Prefix:   p.Prefix,
		IsActive: p.IsActive,
	}

	result := pr.Store.DB.Table("projects").Create(&project).Scan(&p)

	return result.Error
}

func (pr *ProjectRepository) FindAll(filter *utils.Filter) (*[]models.Project, error) {
	project := &[]models.Project{}
	filter = utils.GetDefaultsFilter(filter)

	result := pr.Store.DB.
		Table("projects").
		Find(&project).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		Scan(&project)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return nil, store.ErrRecordNotFound
	}

	return project, result.Error
}

func (pr *ProjectRepository) FindById(id string) (*models.Project, error) {
	project := &models.Project{}

	result := pr.Store.DB.Table("projects").Where("id = ?", id).First(project).Scan(&project)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return project, result.Error
}

func (pr *ProjectRepository) Update(project *models.Project) error {
	result := pr.Store.DB.Table("projects").Save(project).Scan(&project)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return store.ErrRecordNotFound
	}

	return result.Error
}

func (pr *ProjectRepository) Delete(project *models.Project) error {
	result := pr.Store.DB.Table("projects").Delete(project)
	if result.Error != nil {
		return store.ErrRecordNotFound
	}

	return result.Error
}
