package rest

import (
	"github.com/aetherteam/logger_center/internal/config"
	"github.com/aetherteam/logger_center/internal/enums"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/services"
	"github.com/aetherteam/logger_center/internal/utils"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"net/http"
	"time"
)

type UserHandler struct {
	UserService services.UserService
}

type updateUserDTO struct {
	Email string `json:"email"`
}

func (uuDTO *updateUserDTO) Validate() error {
	return validation.ValidateStruct(
		uuDTO,
		validation.Field(&uuDTO.Email, validation.Required, is.Email),
	)
}

func (uh *UserHandler) FindAll(ctx *gin.Context) {
	users, err := uh.UserService.FindAll(&utils.Filter{})
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusNotFound, err.Error(), err)
		return
	}

	var sanitizedUsers []models.User

	for _, user := range *users {
		user.Sanitize()
		sanitizedUsers = append(sanitizedUsers, user)
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResUsersFound, sanitizedUsers)
	return
}

func (uh *UserHandler) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := uh.UserService.FindById(id)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusNotFound, err.Error(), err)
		return
	}

	user.Sanitize()

	utils.ResponseHandler(ctx, http.StatusOK, config.ResUserFound, user)
	return
}

func (uh *UserHandler) Update(ctx *gin.Context) {
	var body updateUserDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest, nil)
		return
	}

	if err := body.Validate(); err != nil {
		splitErr, _ := err.(validation.Errors)
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest, splitErr)
		return
	}

	id := ctx.Param("id")
	user := ctx.Value("user").(*models.User)

	if id != user.ID && (user.Role != enums.Admin.String() || user.Role != enums.Moderator.String()) {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, config.ErrForbiddenAccess, nil)
		return
	}

	updatedUser := models.User{
		ID:        id,
		Email:     body.Email,
		UpdatedAt: time.Now(),
	}

	result, err := uh.UserService.Update(id, &updatedUser)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err.Error(), err)
		return
	}

	result.Sanitize()

	utils.ResponseHandler(ctx, http.StatusOK, config.ResUserUpdated, result)
	return
}

func (uh *UserHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	user := ctx.Value("user").(*models.User)
	if id != user.ID && (user.Role != enums.Admin.String() || user.Role != enums.Moderator.String()) {
		utils.ErrorResponseHandler(ctx, http.StatusForbidden, config.ErrForbiddenAccess, nil)
		return
	}

	if err := uh.UserService.Delete(id); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err.Error(), err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResUserDeleted, nil)
	return
}

func NewUserHandler(us services.UserService) UserHandler {
	return UserHandler{
		UserService: us,
	}
}
