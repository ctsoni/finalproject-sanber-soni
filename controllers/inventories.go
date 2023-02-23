package controllers

import (
	"finalproject-sanber-soni/entity"
	"finalproject-sanber-soni/helper"
	"finalproject-sanber-soni/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type InventoryHandler struct {
	service service.InventoryService
}

func NewInventoryHandler(inventoryService service.InventoryService) *InventoryHandler {
	return &InventoryHandler{inventoryService}
}

func (h *InventoryHandler) InsertInventory(ctx *gin.Context) {
	var inputInventory entity.InputInventory
	currentUser := ctx.MustGet("currentUser").(entity.Users)

	err := ctx.ShouldBindJSON(&inputInventory)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(
			"Insert inventory failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	inventory, stock, err := h.service.AddItem(inputInventory, currentUser.IsAdmin)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Insert inventory failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	msg := helper.FormatInventorySaveResponse(inventory, stock)
	response := helper.APIResponse("Insert success", http.StatusOK, "success", msg)
	ctx.JSON(http.StatusOK, response)
}

func (h *InventoryHandler) UpdateInventory(ctx *gin.Context) {
	var input entity.UpdateInventory
	currentUser := ctx.MustGet("currentUser").(entity.Users)
	id, err := strconv.Atoi(ctx.Param("inven_id"))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Update inventory failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(
			"Update inventory failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	inventory, stock, err := h.service.UpdateItem(input, currentUser.IsAdmin, id)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Update inventory failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	msg := helper.FormatInventorySaveResponse(inventory, stock)
	response := helper.APIResponse("Update success", http.StatusOK, "success", msg)
	ctx.JSON(http.StatusOK, response)
}

func (h *InventoryHandler) DeleteInventory(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(entity.Users)
	id, err := strconv.Atoi(ctx.Param("inven_id"))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Delete inventory failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.service.DeleteItem(currentUser.IsAdmin, id)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Delete inventory failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete success", http.StatusOK, "success", nil)
	ctx.JSON(http.StatusOK, response)
}

func (h *InventoryHandler) GetAll(ctx *gin.Context) {
	inventories, err := h.service.GetAll()
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Get all inventory failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get success", http.StatusOK, "success", inventories)
	ctx.JSON(http.StatusOK, response)
}

func (h *InventoryHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("inven_id"))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Get inventory failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	inventory, err := h.service.GetById(id)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Get inventory failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get success", http.StatusOK, "success", inventory)
	ctx.JSON(http.StatusOK, response)
}
