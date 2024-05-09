package rest

import (
	"github.com/aetherteam/logger_center/internal/config"
	"github.com/aetherteam/logger_center/internal/enums"
	"github.com/aetherteam/logger_center/internal/models"
	"github.com/aetherteam/logger_center/internal/services"
	"github.com/aetherteam/logger_center/internal/utils"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"
	"time"
)

type LogHandler struct {
	LogService services.LogService
}

type updateLogDTO struct {
	ChainID   string         `json:"chain_id"`
	ProjectID string         `json:"project_id"`
	Content   string         `json:"content"`
	Timestamp int64          `json:"Timestamp"`
	Level     enums.LogLevel `json:"level"`
}

type createLogDTO struct {
	ChainID   string         `json:"chain_id"`
	Content   string         `json:"content"`
	Timestamp int64          `json:"Timestamp"`
	Level     enums.LogLevel `json:"level"`
}

func (clDTO *createLogDTO) Validate() error {
	return validation.ValidateStruct(
		clDTO,
		validation.Field(&clDTO.Level, validation.Required),
	)
}

func (lh *LogHandler) Search(ctx *gin.Context) {
	projectID := ctx.Param("project_id")
	if projectID == "" {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	filter := utils.GetDefaultsFilterFromQuery(ctx)

	queryString := ctx.Query("search")
	if len(queryString) < 3 {
		return
	}

	logs, _ := lh.LogService.Search(projectID, queryString, filter)

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogsFound, logs)
	return
}

func (lh *LogHandler) FindAll(ctx *gin.Context) {
	projectID := ctx.Param("project_id")
	if projectID == "" {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}
	filter := utils.GetDefaultsFilterFromQuery(ctx)

	logs, _ := lh.LogService.FindAll(projectID, filter)

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogsFound, logs)
	return
}

func (lh *LogHandler) FindById(ctx *gin.Context) {
	logID := ctx.Param("log_id")

	log, err := lh.LogService.FindById(logID)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusNotFound, err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogFound, log)
	return
}

func (lh *LogHandler) FindByChainId(ctx *gin.Context) {
	chainID := ctx.Param("chain_id")

	filter := utils.GetDefaultsFilterFromQuery(ctx)

	logs, _ := lh.LogService.FindByChainId(chainID, filter)

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogFound, logs)
	return
}

func (lh *LogHandler) Create(ctx *gin.Context) {
	var body createLogDTO

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

	if body.Timestamp <= 0 {
		now := time.Now()
		body.Timestamp = now.Unix()
	}

	log := models.Log{
		ChainID:   body.ChainID,
		ProjectID: projectID,
		Content:   body.Content,
		Timestamp: body.Timestamp,
		Level:     body.Level.String(),
	}

	result, err := lh.LogService.Create(&log)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogCreated, result)
	return
}

func (lh *LogHandler) Update(ctx *gin.Context) {
	var body updateLogDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	id := ctx.Param("log_id")

	if body.Timestamp <= 0 {
		now := time.Now()
		body.Timestamp = now.Unix()
	}

	updatedLog := models.Log{
		ID:        id,
		ChainID:   body.ChainID,
		ProjectID: body.ProjectID,
		Content:   body.Content,
		Timestamp: body.Timestamp,
		Level:     body.Level.String(),
	}

	result, err := lh.LogService.Update(id, &updatedLog)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogUpdated, result)
	return
}

func (lh *LogHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("log_id")

	if err := lh.LogService.Delete(id); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err)
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
