package controllers

import (
	"finalproject-sanber-soni/entity"
	"finalproject-sanber-soni/helper"
	"finalproject-sanber-soni/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ReviewHandler struct {
	service service.ReviewService
}

func NewReviewHandler(service service.ReviewService) *ReviewHandler {
	return &ReviewHandler{service}
}

func (h *ReviewHandler) AddReview(ctx *gin.Context) {
	var input entity.InputReview
	currentUser := ctx.MustGet("currentUser").(entity.Users)

	transID, err := strconv.Atoi(ctx.Param("trans_id"))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Create review failed",
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
			"Create review failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	item, err := h.service.FindItem(transID)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Create review failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	review, err := h.service.AddReview(input, currentUser.ID, transID, currentUser.IsAdmin)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Create review failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	msg := helper.FormatReviewResponse(review, currentUser.FullName, item)
	response := helper.APIResponse("Create review success", http.StatusOK, "success", msg)
	ctx.JSON(http.StatusOK, response)
}

func (h *ReviewHandler) EditReview(ctx *gin.Context) {
	var input entity.InputReview
	currentUser := ctx.MustGet("currentUser").(entity.Users)

	reviewID, err := strconv.Atoi(ctx.Param("review_id"))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Edit review failed",
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
			"Edit review failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	review, err := h.service.EditReview(input, reviewID, currentUser.ID)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Edit review failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	item, err := h.service.FindItem(review.TransId)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Create review failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	msg := helper.FormatReviewResponse(review, currentUser.FullName, item)
	response := helper.APIResponse("Edit review success", http.StatusOK, "success", msg)
	ctx.JSON(http.StatusOK, response)
}

func (h *ReviewHandler) DeleteReview(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(entity.Users)

	reviewID, err := strconv.Atoi(ctx.Param("review_id"))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Delete review failed",
			http.StatusUnprocessableEntity,
			"error",
			errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.service.DeleteReview(reviewID, currentUser.ID)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Delete review failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	msg := gin.H{"message": "delete success"}
	response := helper.APIResponse("Delete review success", http.StatusOK, "success", msg)
	ctx.JSON(http.StatusOK, response)
}

func (h *ReviewHandler) GetAll(ctx *gin.Context) {
	reviews, err := h.service.GetAll()
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Get all reviews failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get all reviews success", http.StatusOK, "success", reviews)
	ctx.JSON(http.StatusOK, response)
}

func (h *ReviewHandler) GetByUserID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Get reviews by user id failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	reviews, err := h.service.GetByUserId(userID)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Get reviews by user id failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get reviews by user id success", http.StatusOK, "success", reviews)
	ctx.JSON(http.StatusOK, response)
}

func (h *ReviewHandler) GetByInvenID(ctx *gin.Context) {
	invenID, err := strconv.Atoi(ctx.Param("inven_id"))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Get reviews by inventory id failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	reviews, err := h.service.GetByInvenId(invenID)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIResponse(
			"Get reviews by inventory id failed",
			http.StatusBadRequest,
			"error",
			errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get reviews by inventory id success", http.StatusOK, "success", reviews)
	ctx.JSON(http.StatusOK, response)
}
