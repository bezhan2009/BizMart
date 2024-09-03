package Category

import (
	"BizMart/models"
	"BizMart/pkg/controllers/handlers"
	"BizMart/pkg/repository/categoryRepository"
	"BizMart/pkg/service/CategoryService"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllCategories retrieves all categories from the repository
func GetAllCategories(c *gin.Context) {
	categories, err := categoryRepository.GetAllCategories() // Corrected package name to lowercase 'category'
	if err != nil {
		handlers.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// GetCategoryById retrieves a category by its ID from the repository
func GetCategoryById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handlers.HandleError(c, err)
		return
	}

	categ, err := categoryRepository.GetCategoryByID(uint(id))
	if err != nil {
		handlers.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categ})
}

// CreateCategory creates a new category in the repository
func CreateCategory(c *gin.Context) {
	var categ models.Category
	if err := c.ShouldBindJSON(&categ); err != nil {
		handlers.HandleError(c, err)
		return
	}

	categoryID, err := CategoryService.CreateCategory(categ)
	if err != nil {
		handlers.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "category created successfully",
		"category_id": categoryID,
	})
}

// UpdateCategory updates an existing category in the repository
func UpdateCategory(c *gin.Context) {
	var categ models.Category
	if err := c.ShouldBindJSON(&categ); err != nil {
		handlers.HandleError(c, err)
		return
	}

	categoryID, err := categoryRepository.UpdateCategory(categ)
	if err != nil {
		handlers.HandleError(c, err)
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
		handlers.HandleError(c, err)
		return
	}

	err = categoryRepository.DeleteCategory(uint(id))
	if err != nil {
		handlers.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category deleted successfully",
	})
}
