package controllers

import (
	"BizMart/internal/app/models"
	"BizMart/internal/app/service"
	"BizMart/internal/controllers/middlewares"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetStores godoc
// @Summary Get all stores
// @Description Fetches all available stores.
// @Tags stores
// @Accept  json
// @Produce  json
// @Success 200 {object} models.DefaultResponse "Returns a list of stores"
// @Failure 400 {object} models.ErrorResponse
// @Router /store [get]
func GetStores(c *gin.Context) {
	stores, err := repository.GetStores()
	if err != nil {
		HandleError(c, err)
	}

	c.JSON(http.StatusOK, gin.H{"stores": stores})
}

// GetStoreByID godoc
// @Summary Get store by ID
// @Description Retrieves a store by its ID along with the number of products and orders.
// @Tags stores
// @Accept  json
// @Produce  json
// @Param id path int true "Store ID"
// @Success 200 {object} models.DefaultResponse "Returns the store, products, product count, and order count"
// @Failure 404 {object} models.ErrorResponse
// @Router /store/{id} [get]
func GetStoreByID(c *gin.Context) {
	storeStrID := c.Param("id")
	storeID, err := strconv.Atoi(storeStrID)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	OurStore, err := service.GetStoreByID(uint(storeID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrStoreNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	storeProducts, err := repository.GetProductByStoreID(uint(storeID))

	productNums, err := repository.GetNumberOfStoreProducts(uint(storeID))
	if err != nil {
		HandleError(c, err)
	}

	orderNums, err := repository.GetNumberOfStoreOrders(uint(storeID))
	if err != nil {
		HandleError(c, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"store":       OurStore,
		"products":    storeProducts,
		"product_num": productNums,
		"order_num":   orderNums,
	})
}

// CreateStore godoc
// @Summary Create a new store
// @Description Creates a new store for the current user.
// @Tags stores
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param store body models.StoreRequest true "Store data"
// @Success 200 {object} models.DefaultResponse "Returns success message"
// @Failure 400 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse "Permission denied"
// @Router /store [post]
func CreateStore(c *gin.Context) {
	var OurStore models.Store
	if err := c.ShouldBindJSON(&OurStore); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	accounts, err := repository.GetAccountsByUserID(userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	if len(accounts) <= 0 {
		HandleError(c, errs.ErrAccountNotFound)
		return
	}

	OurStore.OwnerID = userID

	if err := service.CreateStore(OurStore); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "store created successfully"})
}

// UpdateStore godoc
// @Summary Update an existing store
// @Description Updates the details of a store, such as its name, description, and other properties.
// @Tags stores
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path int true "Store ID"
// @Param store body models.StoreRequest true "Updated store data"
// @Success 200 {object} models.DefaultResponse "Returns success message"
// @Failure 400 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse "Permission denied"
// @Router /store/{id} [put]
func UpdateStore(c *gin.Context) {
	storeStrID := c.Param("id")
	storeID, err := strconv.Atoi(storeStrID)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	var OurStore models.Store
	OurStore.ID = uint(storeID)

	if err = c.ShouldBindJSON(&OurStore); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	store, err := service.GetStoreByID(uint(storeID))
	if err != nil {
		HandleError(c, err)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)

	if userID != store.OwnerID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	OurStore.ID = uint(storeID)
	err = repository.UpdateStore(uint(storeID), &OurStore)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "store updated successfully"})
}

// DeleteStore godoc
// @Summary Delete a store
// @Description Deletes a store by its ID.
// @Tags stores
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path int true "Store ID"
// @Success 200 {object} models.DefaultResponse "Returns success message"
// @Failure 404 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse "Permission denied"
// @Router /store/{id} [delete]
func DeleteStore(c *gin.Context) {
	storeStrID := c.Param("id")
	storeID, err := strconv.Atoi(storeStrID)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	store, err := service.GetStoreByID(uint(storeID))
	if err != nil {
		HandleError(c, err)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)

	if userID != store.OwnerID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	err = repository.DeleteStore(uint(storeID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "store deleted successfully"})
}
