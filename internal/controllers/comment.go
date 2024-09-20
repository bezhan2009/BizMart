package controllers

import (
	"BizMart/internal/app/models"
	"BizMart/internal/controllers/middlewares"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetProductComments godoc
// @Summary Get all comments for a product
// @Description Fetches all comments for a product and builds comment tree.
// @Tags comments
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Comment
// @Failure 404 {object} models.ErrorResponse
// @Router /product/comments/{id} [get]
func GetProductComments(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	// Получаем комментарии к продукту
	mainComments, commentsDict, err := repository.GetProductComments(uint(productID))
	if err != nil {
		HandleError(c, err)
		return
	}

	// Строим дерево комментариев
	commentTree := repository.BuildCommentTree(mainComments, commentsDict)

	c.JSON(http.StatusOK, gin.H{"comments": commentTree})
}

// CreateProductComment godoc
// @Summary Create a comment for a product
// @Description Creates a new comment for a product.
// @Tags comments
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param comment body models.CommentRequest true "Comment data"
// @Success 200 {object} models.DefaultResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /product/comments/{id} [post]
func CreateProductComment(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var commentReq models.CommentRequest
	if err := c.ShouldBindJSON(&commentReq); err != nil {
		HandleError(c, errs.ErrInvalidData)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	// Создаем комментарий через сервис
	_, err = repository.CreateComment(uint(productID), userID, commentReq)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "comment created successfully"})
}

// DeleteComment godoc
// @Summary Delete a comment by ID
// @Description Deletes a comment and its children recursively.
// @Tags comments
// @Security ApiKeyAuth
// @Param id path int true "Comment ID"
// @Success 200 {object} models.DefaultResponse "Comment deleted successfully"
// @Failure 404 {object} models.ErrorResponse
// @Router /product/comments/{id} [delete]
func DeleteComment(c *gin.Context) {
	commentIDStr := c.Param("id")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	err = repository.DeleteComment(uint(commentID), userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
