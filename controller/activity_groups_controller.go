package controller

import (
	"api_activity/services"

	"github.com/gin-gonic/gin"
)

type ActivityGroupController interface {
	FindByID(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	UpdateByID(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

type activityGroupController struct {
	service services.ActivityGroupsService
}

func New(service services.ActivityGroupsService) ActivityGroupController {
	return &activityGroupController{service}
}

func (c *activityGroupController) FindByID(ctx *gin.Context) {

}
func (c *activityGroupController) FindAll(ctx *gin.Context) {

}
func (c *activityGroupController) Create(ctx *gin.Context) {

}
func (c *activityGroupController) UpdateByID(ctx *gin.Context) {

}
func (c *activityGroupController) DeleteByID(ctx *gin.Context) {

}
