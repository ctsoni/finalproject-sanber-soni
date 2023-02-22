package controllers

import (
	"finalproject-sanber-soni/helper"
	"finalproject-sanber-soni/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type categoriesHandler struct {
	catService service.CategoriesService
}

func NewCatHandler(categoriesService service.CategoriesService) *categoriesHandler {
	return &categoriesHandler{categoriesService}
}

func (h *categoriesHandler) GetAllCategories(ctx *gin.Context) {
	categories, err := h.catService.GetAll()
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Get all categories failed",
			http.StatusUnauthorized,
			"error",
			errorMessage)
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	msg := helper.FormatCatGetAllResponse(categories)
	response := helper.APIResponse("Get all users success", http.StatusOK, "success", msg)
	ctx.JSON(http.StatusOK, response)
}
