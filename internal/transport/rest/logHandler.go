package rest

import (
	"github.com/aetherteam/logger_center/internal/config"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/services"
	"github.com/aetherteam/logger_center/internal/utils"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"net/http"
)

type LogHandler struct {
	LogService services.LogService
}

type updateLogDTO struct {
	ChainID   string `json:"chain_id"`
	ProjectID string `json:"project_id"`
	Content   string `json:"content"`
	Level     string `json:"level"`
}

type createLogDTO struct {
	ChainID   string `json:"chain_id"`
	ProjectID string `json:"project_id"`
	Content   string `json:"content"`
	Level     string `json:"level"`
}

func (clDTO *createLogDTO) Validate() error {
	return validation.ValidateStruct(
		clDTO,
		validation.Field(&clDTO.ProjectID, validation.Required, is.UUIDv4),
		validation.Field(&clDTO.Level, validation.Required),
	)
}

func (lh *LogHandler) FindAll(ctx *gin.Context) {
	logs, err := lh.LogService.FindAll(&utils.Filter{})
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusNotFound, err.Error(), err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogsFound, logs)
	return
}

func (lh *LogHandler) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	log, err := lh.LogService.FindById(id)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusNotFound, err.Error(), err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogFound, log)
	return
}

func (lh *LogHandler) FindByChainId(ctx *gin.Context) {
	id := ctx.Param("id")

	logs, err := lh.LogService.FindByChainId(id)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusNotFound, err.Error(), err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogFound, logs)
	return
}

func (lh *LogHandler) Create(ctx *gin.Context) {
	var body createLogDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest, nil)
		return
	}

	if err := body.Validate(); err != nil {
		splitErr, _ := err.(validation.Errors)
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest, splitErr)
		return
	}

	log := models.Log{
		ChainID:   body.ChainID,
		ProjectID: body.ProjectID,
		Content:   body.Content,
		Level:     body.Level,
	}

	result, err := lh.LogService.Create(&log)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err.Error(), err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogCreated, result)
	return
}

func (lh *LogHandler) Update(ctx *gin.Context) {
	var body updateLogDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest, nil)
		return
	}

	id := ctx.Param("id")

	updatedLog := models.Log{
		ID:        id,
		ChainID:   body.ChainID,
		ProjectID: body.ProjectID,
		Content:   body.Content,
		Level:     body.Level,
	}

	result, err := lh.LogService.Update(id, &updatedLog)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err.Error(), err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogUpdated, result)
	return
}

func (lh *LogHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := lh.LogService.Delete(id); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err.Error(), err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogDeleted, nil)
	return
}

func NewLogHandler(ls services.LogService) LogHandler {
	return LogHandler{
		LogService: ls,
	}
}
