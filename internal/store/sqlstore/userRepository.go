package sqlstore

import (
	"errors"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/store"
	"github.com/aetherteam/logger_center/internal/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	Store *Store
}

func (ur *UserRepository) Create(u *models.User) error {
	id := uuid.NewV4().String()
	user := models.User{
		ID:       id,
		Email:    u.Email,
		Role:     u.Role,
		Password: u.Password,
	}

	result := ur.Store.DB.Table("users").Create(&user).Scan(&u)

	return result.Error
}

func (ur *UserRepository) FindAll(filter *utils.Filter) (*[]models.User, error) {
	user := &[]models.User{}

	if filter.Limit == 0 {
		filter.Limit = 10
	}
	if filter.Order == "" {
		filter.Order = "id desc"
	}

	result := ur.Store.DB.
		Table("users").
		Find(&user).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		Scan(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return nil, store.ErrRecordNotFound
	}

	return user, result.Error
}

func (ur *UserRepository) FindById(id string) (*models.User, error) {
	user := &models.User{}

	result := ur.Store.DB.Table("users").Where("id = ?", id).First(user).Scan(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return user, result.Error
}

func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}

	result := ur.Store.DB.Table("users").Where("email = ?", email).First(user).Scan(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return user, result.Error
}

func (ur *UserRepository) Update(user *models.User) error {
	result := ur.Store.DB.Table("users").Save(user).Scan(&user)
	if result.Error != nil {
		return store.ErrRecordNotUpdated
	}

	return result.Error
}

func (ur *UserRepository) Delete(user *models.User) error {
	result := ur.Store.DB.Table("users").Delete(user)
	if result.Error != nil {
		return store.ErrRecordNotDeleted
	}

	return result.Error
}
