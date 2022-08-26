package controller

import (
	"api_activity/requestdata"
	"api_activity/response"
	"api_activity/services"
	"fmt"
	"net/http"
	"strconv"

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
	paramid := ctx.Param("id")
	id, _ := strconv.ParseInt(paramid, 10, 64)
	res, err := c.service.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseRespone{
			Status:  "Not Found",
			Message: fmt.Sprintf("Activity with ID %s Not Found", id),
			Data:    response.EmptyObj{},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.BaseRespone{
		Status:  "Success",
		Message: "Success",
		Data:    res,
	})

}
func (c *activityGroupController) FindAll(ctx *gin.Context) {
	res, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseRespone{
			Status:  err.Error(),
			Message: err.Error(),
			Data:    response.EmptyObj{},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.BaseRespone{
		Status:  "Success",
		Message: "Success",
		Data:    res,
	})
}
func (c *activityGroupController) Create(ctx *gin.Context) {
	var input requestdata.CreateActicityGroups
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseRespone{
			Status:  err.Error(),
			Message: err.Error(),
			Data:    response.EmptyObj{},
		})
		return
	}
	errval := input.Validate()
	if errval != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseRespone{
			Status:  "Bad Request",
			Message: errval.Error(),
			Data:    response.EmptyObj{},
		})
		return
	}
	res, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseRespone{
			Status:  err.Error(),
			Message: err.Error(),
			Data:    response.EmptyObj{},
		})
		return
	}
	ctx.JSON(http.StatusOK, response.BaseRespone{
		Status:  "Success",
		Message: "Success",
		Data:    res,
	})
}
func (c *activityGroupController) UpdateByID(ctx *gin.Context) {
	var input requestdata.UpdateActivityGroups
	paramid := ctx.Param("id")
	id, _ := strconv.ParseInt(paramid, 10, 64)
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseRespone{
			Status:  err.Error(),
			Message: err.Error(),
			Data:    response.EmptyObj{},
		})
		return
	}
	errval := input.Validate()
	if errval != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseRespone{
			Status:  "Bad Request",
			Message: errval.Error(),
			Data:    response.EmptyObj{},
		})
		return
	}
	res, err := c.service.UpdateByID(input, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseRespone{
			Status:  err.Error(),
			Message: err.Error(),
			Data:    response.EmptyObj{},
		})
		return
	}
	ctx.JSON(http.StatusOK, response.BaseRespone{
		Status:  "Success",
		Message: "Success",
		Data:    res,
	})
}
func (c *activityGroupController) DeleteByID(ctx *gin.Context) {
	paramid := ctx.Param("id")
	id, _ := strconv.ParseInt(paramid, 10, 64)
	err := c.service.DeleteByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseRespone{
			Status:  err.Error(),
			Message: err.Error(),
			Data:    response.EmptyObj{},
		})
		return
	}
	ctx.JSON(http.StatusOK, response.BaseRespone{
		Status:  "Success",
		Message: "Success",
		Data:    response.EmptyObj{},
	})
}
