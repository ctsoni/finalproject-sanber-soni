package controllers

import (
	"finalproject-sanber-soni/auth"
	"finalproject-sanber-soni/entity"
	"finalproject-sanber-soni/helper"
	"finalproject-sanber-soni/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// userHandler is object that has userService field with type of service.userService interface contract
type userHandler struct {
	userService service.UserService
}

// NewUserHandler is function to create new userHandler object
func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

// RegisterUser is userHandler method to serve as gin.HandlerFunc for registering user endpoint
func (h *userHandler) RegisterUser(ctx *gin.Context) {
	var input entity.InputRegisterUsers

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(
			"Register failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Register failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := auth.GenerateToken(newUser.ID, newUser.IsAdmin)
	if err != nil {
		response := helper.APIResponse(
			"Register failed",
			http.StatusBadRequest,
			"error",
			"Failed Generate Token")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	format := helper.FormatUserResponse(newUser, token)
	response := helper.APIResponse("Register success", http.StatusOK, "success", format)
	ctx.JSON(http.StatusOK, response)
}

// Login is userHandler method to serve as gin.HandlerFunc for login user endpoint
func (h *userHandler) Login(ctx *gin.Context) {
	var input entity.InputLogin

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(
			"Login failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Login failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := auth.GenerateToken(user.ID, user.IsAdmin)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Login failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	format := helper.FormatUserResponse(user, token)
	response := helper.APIResponse("Login success", http.StatusOK, "success", format)
	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	var input entity.InputUpdateUser

	userId, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Update failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(
			"Update failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedUser, err := h.userService.UpdateUser(userId, input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Update failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	msg := helper.FormatUserEditResponse(updatedUser)
	response := helper.APIResponse("Update success", http.StatusOK, "success", msg)
	ctx.JSON(http.StatusOK, response)
}
