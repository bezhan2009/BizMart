package controllers

import (
	"BizMart/internal/app/models"
	"BizMart/internal/app/service"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllCategories retrieves all categories from the repository
func GetAllCategories(c *gin.Context) {
	categories, err := repository.GetAllCategories() // Corrected package name to lowercase 'category'
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// GetCategoryById retrieves a category by its ID from the repository
func GetCategoryById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, err)
		return
	}

	categ, err := repository.GetCategoryByID(uint(id))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categ})
}

// CreateCategory creates a new category in the repository
func CreateCategory(c *gin.Context) {
	var categ models.Category
	if err := c.ShouldBindJSON(&categ); err != nil {
		HandleError(c, err)
		return
	}

	categoryID, err := service.CreateCategory(categ)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "category created successfully",
		"category_id": categoryID,
	})
}

// UpdateCategory updates an existing category in the repository
func UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, errs.ErrPathParametrized)
		return
	}

	var categ models.Category
	if err := c.ShouldBindJSON(&categ); err != nil {
		HandleError(c, err)
		return
	}

	categoryID, err := service.UpdateCategory(uint(id), categ)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "category updated successfully",
		"category_id": categoryID,
	})
}

// DeleteCategory deletes a category by its ID from the repository
func DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, err)
		return
	}

	err = repository.DeleteCategory(uint(id))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category deleted successfully",
	})
}
