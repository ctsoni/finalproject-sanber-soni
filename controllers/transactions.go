package controllers

import (
	"finalproject-sanber-soni/entity"
	"finalproject-sanber-soni/helper"
	"finalproject-sanber-soni/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(service service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service}
}

func (h *TransactionHandler) CreateTransaction(ctx *gin.Context) {
	var input entity.InputTransaction
	currentUser := ctx.MustGet("currentUser").(entity.Users)

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(
			"Making transaction failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transaction, err := h.service.CreateTransaction(input, currentUser.ID, currentUser.IsAdmin)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Making transaction failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	expMsg := "You have 5 minutes to complete the transaction"
	msg := helper.FormatTransactionResponse(transaction, input.Item, expMsg)
	response := helper.APIResponse("Making transaction success", http.StatusOK, "success", msg)
	ctx.JSON(http.StatusOK, response)

}

func (h *TransactionHandler) UpdateTransaction(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(entity.Users)
	transactionId, err := strconv.Atoi(ctx.Param("trans_id"))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Update transaction failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	status := ctx.Query("action")
	transaction, item, err := h.service.UpdateTransaction(transactionId, status, currentUser.ID, currentUser.IsAdmin)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Update transaction failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	msg := helper.FormatTransactionResponse(transaction, item, "updated")
	response := helper.APIResponse("Update transaction success", http.StatusOK, "success", msg)
	ctx.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) GetAll(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(entity.Users)

	transactions, err := h.service.GetAll(currentUser.ID)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Get all transactions failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get all success", http.StatusOK, "success", transactions)
	ctx.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) GetByStatus(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(entity.Users)
	status := ctx.Query("status")
	err := h.service.ValidateStatus(status)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Get by status transactions failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, err := h.service.GetByStatus(currentUser.ID, status)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Get by status transactions failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get by status success", http.StatusOK, "success", transactions)
	ctx.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) UpdateAdmin(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(entity.Users)
	id, err := strconv.Atoi(ctx.Param("trans_id"))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Update transaction by Admin failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.service.TransactionDoneByAdmin(id, currentUser.IsAdmin)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Update transaction by admin failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	reponse := helper.APIResponse("Update by admin success and stock back up", http.StatusOK, "success", nil)
	ctx.JSON(http.StatusOK, reponse)
}
