package rest

import (
	"github.com/eris-apple/logger_center/internal/config"
	"github.com/eris-apple/logger_center/internal/dto"
	"github.com/eris-apple/logger_center/internal/models"
	"github.com/eris-apple/logger_center/internal/services"
	"github.com/eris-apple/logger_center/internal/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

type LogHandler struct {
	LogService services.LogService
}

func (lh *LogHandler) Search(ctx *gin.Context) {
	filter := utils.GetDefaultsFilterFromQuery(ctx)

	projectID := ctx.Param("project_id")
	queryString := ctx.Query("search")
	if len(queryString) < 3 {
		return
	}

	logs, err := lh.LogService.Search(projectID, queryString, filter)
	if err != nil {
		if err.Error() == config.ErrLogsNotFound.Error() {
			utils.ErrorResponseHandler(ctx, http.StatusNotFound, config.ErrLogsNotFound)
			return
		}

		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogsFound, logs)
	return
}

func (lh *LogHandler) FindAll(ctx *gin.Context) {
	filter := utils.GetDefaultsFilterFromQuery(ctx)

	projectID := ctx.Param("project_id")

	logs, err := lh.LogService.FindAll(projectID, filter)
	if err != nil {
		if err.Error() == config.ErrLogsNotFound.Error() {
			utils.ErrorResponseHandler(ctx, http.StatusNotFound, config.ErrLogsNotFound)
			return
		}

		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogsFound, logs)
	return
}

func (lh *LogHandler) FindById(ctx *gin.Context) {
	logID := ctx.Param("log_id")

	log, err := lh.LogService.FindById(logID)
	if err != nil {
		if err.Error() == config.ErrLogNotFound.Error() {
			utils.ErrorResponseHandler(ctx, http.StatusNotFound, config.ErrLogNotFound)
			return
		}

		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogFound, log)
	return
}

func (lh *LogHandler) FindByChainId(ctx *gin.Context) {
	chainID := ctx.Param("chain_id")

	filter := utils.GetDefaultsFilterFromQuery(ctx)

	logs, err := lh.LogService.FindByChainId(chainID, filter)
	if err != nil {
		if err.Error() == config.ErrLogsNotFound.Error() {
			utils.ErrorResponseHandler(ctx, http.StatusNotFound, config.ErrLogsNotFound)
			return
		}

		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogFound, logs)
	return
}

func (lh *LogHandler) Create(ctx *gin.Context) {
	var body dto.CreateLogDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	if err := body.Validate(); err != nil {
		if strings.Contains(err.Error(), "level: cannot be blank") {
			utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrInvalidLogLevel)
			return
		}
		if strings.Contains(err.Error(), "level: must be a valid value") {
			utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrInvalidLogLevel)
			return
		}
		if strings.Contains(err.Error(), "chain_id: must be a valid UUID") {
			utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrInvalidLogChainID)
			return
		}

		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	projectID := ctx.Param("project_id")

	if body.Timestamp <= 0 {
		now := time.Now()
		body.Timestamp = now.Unix()
	}

	l := models.Log{
		ChainID:   body.ChainID,
		ProjectID: projectID,
		Content:   body.Content,
		Timestamp: body.Timestamp,
		Level:     body.Level,
	}

	result, err := lh.LogService.Create(&l)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResLogCreated, result)
	return
}

func (lh *LogHandler) Update(ctx *gin.Context) {
	var body dto.UpdateLogDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	if err := body.Validate(); err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "level: must be a valid value") {
			utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrInvalidLogLevel)
			return
		}
		if strings.Contains(err.Error(), "chain_id: must be a valid UUID") {
			utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrInvalidLogChainID)
			return
		}

		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	id := ctx.Param("log_id")
	log.Print("id")
	log.Print(id)

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
		Level:     body.Level,
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
