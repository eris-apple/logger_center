package services

import (
	"errors"
	"github.com/aetherteam/logger_center/internal/config"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/store"
	"github.com/aetherteam/logger_center/internal/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type UserService struct {
	UserRepository store.UserRepository
}

func (us UserService) FindAll(filter *utils.Filter) (*[]models.User, error) {
	users, err := us.UserRepository.FindAll(filter)
	if err != nil {
		return nil, errors.New(config.ErrUsersNotFound)
	}

	return users, nil
}

func (us UserService) FindById(id string) (*models.User, error) {
	user, err := us.UserRepository.FindById(id)
	if err != nil {
		return nil, errors.New(config.ErrUserNotFound)
	}

	return user, nil
}

func (us UserService) Update(id string, updatedUser *models.User) (*models.User, error) {
	user, _ := us.FindById(id)

	if validation.IsEmpty(updatedUser.Password) {
		updatedUser.Password = user.Password
	}

	if validation.IsEmpty(updatedUser.Email) {
		updatedUser.Email = user.Email
	}

	if validation.IsEmpty(updatedUser.Role) {
		updatedUser.Role = user.Role
	}

	if validation.IsEmpty(updatedUser.Status) {
		updatedUser.Status = user.Status
	}

	updatedUser.UpdatedAt = time.Now()

	if err := us.UserRepository.Update(updatedUser); err != nil {
		return nil, errors.New(config.ErrInternalServerError)
	}

	return updatedUser, nil
}

func (us UserService) Delete(id string) error {
	user, _ := us.FindById(id)

	if err := us.UserRepository.Delete(user); err != nil {
		return errors.New(config.ErrInternalServerError)
	}

	return nil
}

func NewUserService(ur store.UserRepository) UserService {
	return UserService{
		UserRepository: ur,
	}
}
