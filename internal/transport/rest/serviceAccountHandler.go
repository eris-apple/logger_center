package rest

import (
	"github.com/aetherteam/logger_center/internal/services"
	"github.com/gin-gonic/gin"
)

type ServiceAccountHandler struct {
	sah *services.ServiceAccountService
}

func (sah *ServiceAccountHandler) FindAll(ctx *gin.Context) {

}

func (sah *ServiceAccountHandler) FindById(ctx *gin.Context) {

}

func (sah *ServiceAccountHandler) FindBySecret(ctx *gin.Context) {

}

func (sah *ServiceAccountHandler) Create(ctx *gin.Context) {

}

func (sah *ServiceAccountHandler) Update(ctx *gin.Context) {

}

func (sah *ServiceAccountHandler) Delete(ctx *gin.Context) {

}

func NewServiceAccountHandler(sah *services.ServiceAccountService) *ServiceAccountHandler {
	return &ServiceAccountHandler{
		sah: sah,
	}
}
