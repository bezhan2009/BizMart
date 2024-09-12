package controllers

import (
	"BizMart/internal/app/models"
	"BizMart/internal/app/service"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllCategories retrieves all categories from the repository
// @Summary      Get all categories
// @Description  Retrieves a list of all categories
// @Tags         Category
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Category
// @Failure      500  {object}  models.ErrorResponse
// @Router       /category [get]
func GetAllCategories(c *gin.Context) {
	categories, err := repository.GetAllCategories() // Corrected package name to lowercase 'category'
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// GetCategoryById retrieves a category by its ID from the repository
// @Summary      Get category by ID
// @Description  Retrieves a category by its ID
// @Tags         Category
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Success      200  {object}  models.Category
// @Failure      400  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /category/{id} [get]
func GetCategoryById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, err)
		return
	}

	categ, err := repository.GetCategoryByID(uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrCategoryNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categ})
}

// CreateCategory creates a new category in the repository
// @Summary Создать новую категорию
// @Description Создаёт новую категорию с указанными параметрами
// @Tags Category
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param category body models.CategoryRequest true "Данные для новой категории"
// @Success 200 {object} models.DefaultResponse "Успешное создание"
// @Failure 400 {object} models.ErrorResponse "Ошибка в запросе"
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера"
// @Router /category [post]
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
// @Summary Обновить категорию
// @Description Обновляет информацию о категории по её ID
// @Tags Category
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "ID категории"
// @Param category body models.CategoryRequest true "Данные для обновления категории"
// @Success 200 {object} models.DefaultResponse "Успешное обновление"
// @Failure 400 {object} models.ErrorResponse "Ошибка в запросе"
// @Failure 500 {object} models.ErrorResponse "Ошибка сервера"
// @Router /category/{id} [put]
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
// @Summary      Delete category by ID
// @Description  Deletes a category by its ID
// @Tags         Category
// @Security ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Success      200  {object}  models.DefaultResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /category/{id} [delete]
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
