package controller

import (
	"api-crud-golang/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ControllerPostUser godoc
// @Summary Create a new User
// @Description Accepts JSON and creates a new record
// @Tags User
// @Accept json
// @Produce json
// @Param Input body model.UserFill true "User Input"
// @Router /user/post-model [post]
func (c *MainController) CotnrollerPostUser(ctx *gin.Context) {
	var param model.User

	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, model.MetadataResponse{
			Data:    err.Error(),
			Status:  "400",
			Message: "error",
		})
		return
	}

	if err := c.Service.ServicePostUser(&param); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.MetadataResponse{
			Data:    err.Error(),
			Status:  "500",
			Message: "error",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.MetadataResponse{
		Data:    param,
		Status:  "200",
		Message: "succes",
	})
}

// ControllerDeleteUser godoc
// @Summary Delete User
// @Description Hit to delete record
// @Tags User
// @Param id path string true "Main Model Input"
// @Router /user/delete-model/{id} [delete]
func (c *MainController) ControllerDeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.Service.ServiceDeleteUser(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.MetadataResponse{
			Data:    err.Error(),
			Status:  "500",
			Message: "error",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.MetadataResponse{
		Data:    "",
		Status:  "200",
		Message: "succes",
	})
}

// ControllerLogin godoc
// @Summary Login user
// @Description Authenticates user and returns JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param LoginRequest body model.LoginRequest true "Login Payload"
// @Success 200 {object} model.MetadataResponse
// @Failure 400 {object} model.MetadataResponse
// @Failure 401 {object} model.MetadataResponse
// @Router /auth/login [post]
func (c *MainController) ControllerLogin(ctx *gin.Context) {
	var req model.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	// token, err := c.Service.LoginService(req.Email, req.Password);
	tokenStr, err := c.Service.LoginService(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.MetadataResponse{
		Data:    tokenStr,
		Status:  "",
		Message: "succes",
	})
}
