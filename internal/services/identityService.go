package services

import (
	"errors"
	"fmt"
	baseConfig "github.com/aetherteam/logger_center/internal/config"
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

func (is *IdentityService) SignUp(user *models.User) (*models.User, error) {
	password, _ := utils.HashString(user.Password)

	user.Password = password
	user.Role = enums.Guest.String()

	if err := is.UserRepository.Create(user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New(baseConfig.ErrUserAlreadyExist)
		}

		return nil, errors.New(baseConfig.ErrInternalServerError)
	}

	return user, nil
}

func (is *IdentityService) SignIn(credentials *models.User) (*SingInResponse, error) {
	user, err := is.UserRepository.FindByEmail(credentials.Email)
	if err != nil {
		return nil, errors.New(baseConfig.ErrIncorrectEmailOrPassword)
	}

	if ch := utils.CheckHash(credentials.Password, user.Password); ch != true {
		return nil, errors.New(baseConfig.ErrIncorrectEmailOrPassword)
	}

	atTime := time.Now().Add(time.Hour * 24 * 31)

	atClaims := &TokenClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(atTime),
		},
	}

	atJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	atToken, atErr := atJWT.SignedString([]byte(os.Getenv(baseConfig.EnvKeyJWTSecret)))
	if atErr != nil {
		fmt.Print(os.Getenv(baseConfig.EnvKeyJWTSecret))
		fmt.Print(atErr)
		return nil, errors.New("config.ErrInternalServerError2")
	}

	session := &models.Session{
		UserID:   user.ID,
		IsActive: true,
		Token:    atToken,
	}

	if err := is.SessionRepository.Create(session); err != nil {
		return nil, errors.New(baseConfig.ErrInternalServerError)
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
