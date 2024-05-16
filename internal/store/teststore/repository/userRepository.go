package teststore_repository

import (
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/store"
	"github.com/eris-apple/logger_center/internal/utils"
	uuid "github.com/satori/go.uuid"
	"strings"
)

type UserRepository struct {
	Users map[string]*models.User
}

func (ur *UserRepository) Search(filter *utils.Filter, queryString string) ([]models.User, error) {
	filter = utils.GetDefaultsFilter(filter)

	var au []models.User
	for _, u := range ur.Users {
		au = append(au, *u)
	}

	fUsers := utils.FilterArray(au, filter)
	if fUsers == nil {
		return nil, store.ErrRecordNotFound
	}

	var users []models.User
	if queryString != "" {
		for _, u := range fUsers {
			if strings.Contains(u.Email, queryString) {
				users = append(users, u)
			}
		}
	} else {
		users = fUsers
	}

	return fUsers, nil
}

func (ur *UserRepository) FindAll(filter *utils.Filter, where map[string]interface{}) ([]models.User, error) {
	filter = utils.GetDefaultsFilter(filter)

	var au []models.User
	for _, u := range ur.Users {
		au = append(au, *u)
	}

	fUsers := utils.FilterArray(au, filter)
	if fUsers == nil {
		return nil, store.ErrRecordNotFound
	}

	var users []models.User
	if len(where) != 0 {
		for _, u := range fUsers {
			if where["email"] == u.Email {
				users = append(users, u)
			}
			if where["role"] == u.Role {
				users = append(users, u)
			}
			if where["status"] == u.Status {
				users = append(users, u)
			}
		}
	} else {
		users = fUsers
	}

	return fUsers, nil
}

func (ur *UserRepository) FindById(id string) (*models.User, error) {
	user := ur.Users[id]
	if user == nil {
		return nil, store.ErrRecordNotFound
	}

	return user, nil
}

func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	for _, user := range ur.Users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

func (ur *UserRepository) Create(user *models.User) error {
	id := uuid.NewV4().String()
	*user = models.User{
		ID:       id,
		Email:    user.Email,
		Role:     user.Role,
		Status:   user.Status,
		Password: user.Password,
	}

	ur.Users[id] = user

	return nil
}

func (ur *UserRepository) Update(user *models.User) error {
	checkUser := ur.Users[user.ID]
	if checkUser == nil {
		return store.ErrRecordNotFound
	}

	ur.Users[user.ID] = user

	return nil
}

func (ur *UserRepository) Delete(user *models.User) error {
	checkUser := ur.Users[user.ID]
	if checkUser == nil {
		return store.ErrRecordNotFound
	}

	ur.Users[user.ID] = nil

	return nil
}
