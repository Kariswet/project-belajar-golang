package controller

import (
	"api-crud-golang/model"
	"api-crud-golang/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainController struct {
	// Router  *gin.RouterGroup
	Service *service.MainService
}

func NewMainController(service *service.MainService) *MainController {
	return &MainController{
		Service: service,
	}
}

// ControllerPostMainModel godoc
// @Summary Create a new Main Model
// @Description Accepts JSON and creates a new record
// @Tags Admin
// @Accept json
// @Produce json
// @Param MainModel body model.MainModelFill true "Main Modep Input"
// @Security BearerAuth
// @Router /admin/post-model [post]
func (c *MainController) ControllerPostMainModel(ctx *gin.Context) {
	var param model.MainModel

	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.Service.ServicePostMainModel(&param); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, param)
}

// ControllerPostMainModel godoc
// @Summary Get All Main Model
// @Description Hit to get record
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Router /admin/get-model [get]
func (c *MainController) ControllerGetMainModel(ctx *gin.Context) {
	data, err := c.Service.ServiceGetMainModel()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.MetadataResponse{
		Data:    data,
		Status:  "200",
		Message: "succes",
	})
}

// ControllerPostMainModel godoc
// @Summary Get Main Model by ID
// @Description Hit to get record
// @Tags Admin
// @Produce json
// @Param id path string true "Main Modep Input"
// @Security BearerAuth
// @Router /admin/get-by-id-model/{id} [post]
func (c *MainController) ControllerGetByIdMainModel(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := c.Service.ServiceGetByIdMainModel(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.MetadataResponse{
		Data:    data,
		Status:  "200",
		Message: "succes",
	})

}

// ControllerPostMainModel godoc
// @Summary Update a Main Model
// @Description Accepts JSON and update a record
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path string true "Record ID"
// @Param MainModel body model.MainModelFill true "Main Modep Input"
// @Security BearerAuth
// @Router /admin/put-model/{id} [put]
func (c *MainController) ControllerUpdateMainModel(ctx *gin.Context) {
	id := ctx.Param("id")
	var param model.MainModel

	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.Service.ServiceUpdateMainModel(id, &param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	update_data := param

	ctx.JSON(http.StatusOK, model.MetadataResponse{
		Data:    update_data,
		Status:  "200",
		Message: "succes",
	})
}

// ControllerPostMainModel godoc
// @Summary Delete Main Model
// @Description Hit to delete record
// @Tags Admin
// @Param id path string true "Main Model Input"
// @Security BearerAuth
// @Router /admin/delete-model/{id} [delete]
func (c *MainController) ControllerDeleteMainModel(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.Service.ServiceDeleteMainModel(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.MetadataResponse{
			Data:    err.Error(),
			Status:  "500",
			Message: "Found Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.MetadataResponse{
		Data:    "",
		Status:  "200",
		Message: "record succes deleted",
	})
}
