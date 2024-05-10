package repository

import (
	"errors"
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/store"
	"github.com/eris-apple/logger_center/internal/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (ur *UserRepository) Search(filter *utils.Filter, queryString string) (*[]models.User, error) {
	user := &[]models.User{}
	filter = utils.GetDefaultsFilter(filter)

	result := ur.DB.
		Table("users").
		Find(&user).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		Where("email ILIKE ?", "%"+queryString+"%").
		Scan(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return nil, store.ErrRecordNotFound
	}

	return user, result.Error
}

func (ur *UserRepository) FindAll(filter *utils.Filter, where map[string]interface{}) (*[]models.User, error) {
	user := &[]models.User{}
	filter = utils.GetDefaultsFilter(filter)

	result := ur.DB.
		Table("users").
		Find(&user).
		Offset(filter.Offset).
		Limit(filter.Limit).
		Order(filter.Order).
		Where(where).
		Scan(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return nil, store.ErrRecordNotFound
	}

	return user, result.Error
}

func (ur *UserRepository) FindById(id string) (*models.User, error) {
	user := &models.User{}

	result := ur.DB.Table("users").Where("id = ?", id).First(user).Scan(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return user, result.Error
}

func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}

	result := ur.DB.Table("users").Where("email = ?", email).First(user).Scan(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}

	return user, result.Error
}

func (ur *UserRepository) Create(u *models.User) error {
	id := uuid.NewV4().String()
	user := models.User{
		ID:       id,
		Email:    u.Email,
		Role:     u.Role,
		Status:   u.Status,
		Password: u.Password,
	}

	result := ur.DB.Table("users").Create(&user).Scan(&u)
	if result.Error != nil {
		return store.ErrRecordNotCreated
	}

	return result.Error
}

func (ur *UserRepository) Update(user *models.User) error {
	result := ur.DB.Table("users").Save(user).Scan(&user)
	if result.Error != nil {
		return store.ErrRecordNotUpdated
	}

	return result.Error
}

func (ur *UserRepository) Delete(user *models.User) error {
	result := ur.DB.Table("users").Delete(user)
	if result.Error != nil {
		return store.ErrRecordNotDeleted
	}

	return result.Error
}
