package rest

import (
	"github.com/eris-apple/logger_center/internal/config"
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/services"
	"github.com/eris-apple/logger_center/internal/utils"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"net/http"
)

type IdentityHandler struct {
	IdentityService services.IdentityService
}

type signDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (sDTO *signDTO) Validate() error {
	return validation.ValidateStruct(
		sDTO,
		validation.Field(&sDTO.Email, validation.Required, is.Email),
		validation.Field(&sDTO.Password, validation.Required, validation.Length(8, 32)),
	)
}

func (ih *IdentityHandler) SignUp(ctx *gin.Context) {
	var body signDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	if err := body.Validate(); err != nil {
		splitErr, _ := err.(validation.Errors)
		utils.ErrorResponseValidationHandler(ctx, http.StatusBadRequest, config.ErrBadRequest, splitErr)
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
	var body signDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	if err := body.Validate(); err != nil {
		splitErr, _ := err.(validation.Errors)
		utils.ErrorResponseValidationHandler(ctx, http.StatusBadRequest, config.ErrBadRequest, splitErr)
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
