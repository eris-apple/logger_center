package rest

import (
	"github.com/aetherteam/logger_center/internal/config"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/services"
	"github.com/aetherteam/logger_center/internal/utils"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"
	"time"
)

type ServiceAccountHandler struct {
	ServiceAccountService *services.ServiceAccountService
}

type updateAccountServiceDTO struct {
	IsActive bool   `json:"is_active"`
	Name     string `json:"name"`
}

func (uasDTO *updateAccountServiceDTO) Validate() error {
	return validation.ValidateStruct(
		uasDTO,
	)
}

func (sah *ServiceAccountHandler) FindAll(ctx *gin.Context) {
	filter := utils.GetDefaultsFilterFromQuery(ctx)

	sAccounts, _ := sah.ServiceAccountService.FindAll(filter)

	var sanitizedServiceAccounts []models.ServiceAccount

	for _, sAccount := range *sAccounts {
		sAccount.Sanitize()
		sanitizedServiceAccounts = append(sanitizedServiceAccounts, sAccount)
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResServiceAccountsFound, sanitizedServiceAccounts)
	return

}

func (sah *ServiceAccountHandler) FindById(ctx *gin.Context) {
	id := ctx.Param("service_account_id")

	sAccount, err := sah.ServiceAccountService.FindById(id)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusNotFound, err)
		return
	}

	sAccount.Sanitize()

	utils.ResponseHandler(ctx, http.StatusOK, config.ResServiceAccountFound, sAccount)
	return
}

func (sah *ServiceAccountHandler) Create(ctx *gin.Context) {
	var body updateAccountServiceDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	if err := body.Validate(); err != nil {
		splitErr, _ := err.(validation.Errors)
		utils.ErrorResponseValidationHandler(ctx, http.StatusBadRequest, config.ErrBadRequest, splitErr)
		return
	}

	projectID := ctx.Param("project_id")

	log := models.ServiceAccount{
		ProjectID: projectID,
		Name:      body.Name,
		IsActive:  body.IsActive,
		Secret:    utils.RandStringBytes(24),
	}

	result, err := sah.ServiceAccountService.Create(&log)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResServiceAccountCreated, result)
	return

}

func (sah *ServiceAccountHandler) Update(ctx *gin.Context) {
	var body updateAccountServiceDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	if err := body.Validate(); err != nil {
		splitErr, _ := err.(validation.Errors)
		utils.ErrorResponseValidationHandler(ctx, http.StatusBadRequest, config.ErrBadRequest, splitErr)
		return
	}

	id := ctx.Param("service_account_id")

	updatedServiceAccount := models.ServiceAccount{
		ID:        id,
		Name:      body.Name,
		IsActive:  body.IsActive,
		UpdatedAt: time.Now(),
	}

	result, err := sah.ServiceAccountService.Update(id, &updatedServiceAccount)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	result.Sanitize()

	utils.ResponseHandler(ctx, http.StatusOK, config.ResServiceAccountUpdated, result)
	return

}

func (sah *ServiceAccountHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("service_account_id")

	if err := sah.ServiceAccountService.Delete(id); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResServiceAccountDeleted, nil)
	return
}

func NewServiceAccountHandler(sah *services.ServiceAccountService) *ServiceAccountHandler {
	return &ServiceAccountHandler{
		ServiceAccountService: sah,
	}
}
