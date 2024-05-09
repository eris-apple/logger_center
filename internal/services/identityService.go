package services

import (
	"errors"
	"github.com/aetherteam/logger_center/internal/config"
	"github.com/aetherteam/logger_center/internal/enums"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/store"
	"github.com/aetherteam/logger_center/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"os"
	"time"
)

type IdentityService struct {
	UserRepository    store.UserRepository
	SessionRepository store.SessionRepository
}

type TokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type SingInResponse struct {
	*models.User
	AccessToken string `json:"access_token"`
}

func (is *IdentityService) SignUp(user *models.User) (*models.User, *config.APIError) {
	_, err := is.UserRepository.FindByEmail(user.Email)
	if err == nil {
		return nil, config.ErrUserAlreadyExist
	}

	password, _ := utils.HashString(user.Password)

	user.Password = password
	user.Role = enums.Guest.String()
	user.Status = enums.Pending.String()

	if err := is.UserRepository.Create(user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, config.ErrUserAlreadyExist
		}

		return nil, config.ErrInternalServerError
	}

	return user, nil
}

func (is *IdentityService) SignIn(credentials *models.User) (*SingInResponse, *config.APIError) {
	user, err := is.UserRepository.FindByEmail(credentials.Email)
	if err != nil {
		return nil, config.ErrIncorrectEmailOrPassword
	}

	if ch := utils.CheckHash(credentials.Password, user.Password); ch != true {
		return nil, config.ErrIncorrectEmailOrPassword
	}

	atTime := time.Now().Add(time.Hour * 24 * 31)

	atClaims := &TokenClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(atTime),
		},
	}

	atJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	atToken, atErr := atJWT.SignedString([]byte(os.Getenv(config.EnvKeyJWTSecret)))
	if atErr != nil {
		return nil, config.ErrInternalServerError
	}

	session := &models.Session{
		UserID:   user.ID,
		IsActive: true,
		Token:    atToken,
	}

	if err := is.SessionRepository.Create(session); err != nil {
		return nil, config.ErrInternalServerError
	}

	return &SingInResponse{
		User:        user,
		AccessToken: atToken,
	}, nil
}

func NewIdentityService(ur store.UserRepository, sr store.SessionRepository) IdentityService {
	return IdentityService{
		UserRepository:    ur,
		SessionRepository: sr,
	}
}
