package rest

import (
	"github.com/eris-apple/logger_center/internal/config"
	"github.com/eris-apple/logger_center/internal/dto"
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/services"
	"github.com/eris-apple/logger_center/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type IdentityHandler struct {
	IdentityService services.IdentityService
}

func (ih *IdentityHandler) SignUp(ctx *gin.Context) {
	var body dto.SignDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	if err := body.Validate(); err != nil {
		if strings.Contains(err.Error(), "must be a valid email address") {
			utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrInvalidEmail)
			return
		}
		if strings.Contains(err.Error(), "the length must be between 8 and 32") {
			utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrInvalidPassword)
			return
		}

		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	user := &models.User{
		Email:    body.Email,
		Password: body.Password,
	}

	cu, err := ih.IdentityService.SignUp(user)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	cu.Sanitize()

	utils.ResponseHandler(ctx, http.StatusOK, config.ResUserCreated, cu)
	return
}

func (ih *IdentityHandler) SignIn(ctx *gin.Context) {
	var body dto.SignDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	if err := body.Validate(); err != nil {
		if strings.Contains(err.Error(), "must be a valid email address") {
			utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrInvalidEmail)
			return
		}
		if strings.Contains(err.Error(), "the length must be between 8 and 32") {
			utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrInvalidPassword)
			return
		}

		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	credentials := &models.User{Email: body.Email, Password: body.Password}
	user, err := ih.IdentityService.SignIn(credentials)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	user.Sanitize()

	utils.ResponseHandler(ctx, http.StatusOK, config.ReqUserAuthorized, user)
	return
}

func (ih *IdentityHandler) Check(ctx *gin.Context) {
	user := ctx.Value("user").(*models.User)

	utils.ResponseHandler(ctx, http.StatusOK, config.ReqUserAuthorized, user)
	return
}

func NewIdentityHandler(is services.IdentityService) IdentityHandler {
	return IdentityHandler{
		IdentityService: is,
	}
}
