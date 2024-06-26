package rest

import (
	"github.com/eris-apple/logger_center/internal/config"
	"github.com/eris-apple/logger_center/internal/dto"
	"github.com/eris-apple/logger_center/internal/enums"
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/services"
	"github.com/eris-apple/logger_center/internal/utils"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"
	"time"
)

type UserHandler struct {
	UserService services.UserService
}

func (uh *UserHandler) FindAll(ctx *gin.Context) {
	filter := utils.GetDefaultsFilterFromQuery(ctx)

	var payload dto.FindUsersDTO
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, err)
		return
	}

	where, structErr := utils.StructToMap(&payload)
	if structErr != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, structErr)
		return
	}

	users, _ := uh.UserService.FindAll(filter, where)

	var sanitizedUsers []models.User

	if users != nil {
		for _, user := range users {
			user.Sanitize()
			sanitizedUsers = append(sanitizedUsers, user)
		}
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResUsersFound, sanitizedUsers)
	return
}

func (uh *UserHandler) Search(ctx *gin.Context) {
	filter := utils.GetDefaultsFilterFromQuery(ctx)

	queryString := ctx.Query("search")
	if len(queryString) < 3 {
		return
	}

	users, _ := uh.UserService.Search(filter, queryString)

	var sanitizedUsers []models.User

	if users != nil {
		for _, user := range users {
			user.Sanitize()
			sanitizedUsers = append(sanitizedUsers, user)
		}
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResUsersFound, sanitizedUsers)
	return
}

func (uh *UserHandler) FindById(ctx *gin.Context) {
	id := ctx.Param("user_id")

	user, err := uh.UserService.FindById(id)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusNotFound, err)
		return
	}

	user.Sanitize()

	utils.ResponseHandler(ctx, http.StatusOK, config.ResUserFound, user)
	return
}

func (uh *UserHandler) Update(ctx *gin.Context) {
	var body dto.UpdateUserDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	id := ctx.Param("user_id")
	user := ctx.Value("user").(*models.User)
	isUserAdmin := user.Role == enums.Admin.String()
	isUserModerator := user.Role == enums.Moderator.String()

	if id != user.ID && !(isUserAdmin || isUserModerator) {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, config.ErrForbiddenAccess)
		return
	}

	if (!validation.IsEmpty(body.Role) || !validation.IsEmpty(body.Status)) && !(isUserAdmin || isUserModerator) {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, config.ErrForbiddenAccess)
		return
	}

	updatedUser := models.User{
		ID:        id,
		Email:     body.Email,
		UpdatedAt: time.Now(),
	}

	result, err := uh.UserService.Update(id, &updatedUser)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	result.Sanitize()

	utils.ResponseHandler(ctx, http.StatusOK, config.ResUserUpdated, result)
	return
}

func (uh *UserHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("user_id")
	user := ctx.Value("user").(*models.User)
	if id != user.ID && (user.Role != enums.Admin.String() || user.Role != enums.Moderator.String()) {
		utils.ErrorResponseHandler(ctx, http.StatusForbidden, config.ErrForbiddenAccess)
		return
	}

	if err := uh.UserService.Delete(id); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err)
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
