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

type ProjectHandler struct {
	ProjectService services.ProjectService
}

func (ph *ProjectHandler) FindAll(ctx *gin.Context) {
	filter := utils.GetDefaultsFilterFromQuery(ctx)

	var payload dto.FindProjectsDTO
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, err)
		return
	}

	where, structErr := utils.StructToMap(&payload)
	if structErr != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, structErr)
		return
	}

	projects, _ := ph.ProjectService.FindAll(filter, where)

	utils.ResponseHandler(ctx, http.StatusOK, config.ResProjectsFound, projects)
	return
}

func (ph *ProjectHandler) FindById(ctx *gin.Context) {
	id := ctx.Param("project_id")

	project, err := ph.ProjectService.FindById(id)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusNotFound, err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResProjectFound, project)
	return
}

func (ph *ProjectHandler) Create(ctx *gin.Context) {
	var body dto.CreateProjectDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	if err := body.Validate(); err != nil {
		if strings.Contains(err.Error(), "name: cannot be blank") {
			utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrInvalidProjectName)
			return
		}
		if strings.Contains(err.Error(), "prefix: cannot be blank") {
			utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrInvalidProjectPrefix)
			return
		}

		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	project := &models.Project{
		Name:        body.Name,
		Prefix:      body.Prefix,
		IsActive:    body.IsActive,
		Description: body.Description,
	}

	result, err := ph.ProjectService.Create(project)
	if err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResProjectCreated, result)
	return
}

func (ph *ProjectHandler) Update(ctx *gin.Context) {
	var body dto.UpdateProjectDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusBadRequest, config.ErrBadRequest)
		return
	}

	ID := ctx.Param("project_id")
	fResult, FErr := ph.ProjectService.FindById(ID)
	if FErr != nil || fResult == nil {
		utils.ErrorResponseHandler(ctx, http.StatusNotFound, config.ErrProjectNotFound)
		return
	}

	if body.Name == "" {
		body.Name = fResult.Name
	}

	if body.Prefix == "" {
		body.Prefix = fResult.Prefix
	}

	project := models.Project{
		ID:          fResult.ID,
		Name:        body.Name,
		Prefix:      body.Prefix,
		IsActive:    body.IsActive,
		Description: body.Description,
		UpdatedAt:   time.Now(),
	}

	uResult, UErr := ph.ProjectService.Update(&project)
	if UErr != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, UErr)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResProjectUpdated, uResult)
	return
}

func (ph *ProjectHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("project_id")

	if err := ph.ProjectService.Delete(&models.Project{ID: id}); err != nil {
		utils.ErrorResponseHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.ResponseHandler(ctx, http.StatusOK, config.ResProjectDeleted, nil)
	return
}

func NewProjectHandler(ps services.ProjectService) ProjectHandler {
	return ProjectHandler{
		ProjectService: ps,
	}
}
