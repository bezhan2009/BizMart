package controllers

import (
	"BizMart/internal/app/models"
	"BizMart/internal/app/service"
	"BizMart/internal/controllers/middlewares"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetStores(c *gin.Context) {
	stores, err := repository.GetStores()
	if err != nil {
		HandleError(c, err)
	}

	c.JSON(http.StatusOK, gin.H{"stores": stores})
}

func GetStoreByID(c *gin.Context) {
	storeStrID := c.Param("id")
	storeID, err := strconv.Atoi(storeStrID)
	if err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	OurStore, err := service.GetStoreByID(uint(storeID))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"store": OurStore})
}

func CreateStore(c *gin.Context) {
	var OurStore models.Store
	if err := c.ShouldBindJSON(&OurStore); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)

	OurStore.OwnerID = userID

	if err := service.CreateStore(OurStore); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "store created successfully"})
}

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
