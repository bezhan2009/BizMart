package controllers

import (
	"BizMart/models"
	"BizMart/pkg/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllCategories(c *gin.Context) {
	categories, err := repository.GetAllCategories()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func GetCategoryById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(c, err)
		return
	}

	category, err := repository.GetCategoryByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": category})
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		handleError(c, err)
		return
	}

	categoryID, err := repository.CreateCategory(category)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "category created successfully",
		"category_id": categoryID,
	})
}

func UpdateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		handleError(c, err)
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		handleError(c, err)
		return
	}

	categoryID, err := repository.UpdateCategory(category)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "category updated successfully",
		"category_id": categoryID,
	})
}

func DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(c, err)
		return
	}

	err = repository.DeleteCategory(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category deleted successfully",
	})
}
