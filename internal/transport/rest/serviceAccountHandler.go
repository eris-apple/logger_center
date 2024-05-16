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
	"time"
)

type ServiceAccountHandler struct {
	ServiceAccountService *services.ServiceAccountService
}

func (sah *ServiceAccountHandler) Search(ctx *gin.Context) {
	filter := utils.GetDefaultsFilterFromQuery(ctx)
	projectID := ctx.Param("project_id")
	queryString := ctx.Query("search")

	sAccounts, _ := sah.ServiceAccountService.Search(projectID, queryString, filter)

	utils.ResponseHandler(ctx, http.StatusOK, config.ResServiceAccountsFound, sAccounts)
	return
}

func (sah *ServiceAccountHandler) FindAll(ctx *gin.Context) {
	filter := utils.GetDefaultsFilterFromQuery(ctx)
	projectID := ctx.Param("project_id")

	sAccounts, _ := sah.ServiceAccountService.FindAll(projectID, filter)

	utils.ResponseHandler(ctx, http.StatusOK, config.ResServiceAccountsFound, sAccounts)
	return
}

func (sah *ServiceAccountHandler) FindById(ctx *gin.Context) {
	id := ctx.Param("service_account_id")

	sAccount, err := sah.ServiceAccountService.FindById(id)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusNotFound, err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResServiceAccountFound, sAccount)
	return
}

func (sah *ServiceAccountHandler) Create(ctx *gin.Context) {
	var body dto.CreateServiceAccountDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	if err := body.Validate(); err != nil {
		if strings.Contains(err.Error(), "name: cannot be blank") {
			utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrInvalidServiceAccountName)
			return
		}

		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	projectID := ctx.Param("project_id")

	log := models.ServiceAccount{
		ProjectID:   projectID,
		Name:        body.Name,
		IsActive:    body.IsActive,
		Description: body.Description,
		Secret:      utils.RandStringBytes(24),
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
	var body dto.UpdateServiceAccountDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	if err := body.Validate(); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	id := ctx.Param("service_account_id")

	updatedServiceAccount := models.ServiceAccount{
		ID:          id,
		Name:        body.Name,
		IsActive:    body.IsActive,
		Description: body.Description,
		UpdatedAt:   time.Now(),
	}

	result, err := sah.ServiceAccountService.Update(id, &updatedServiceAccount)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err)
		return
	}

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
