package storesControllers

import (
	"BizMart/errs"
	"BizMart/models"
	"BizMart/pkg/controllers/handlers"
	"BizMart/pkg/controllers/middlewares"
	"BizMart/pkg/repository/storeRepository"
	"BizMart/pkg/service/storeService"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetStores(c *gin.Context) {
	stores, err := storeRepository.GetStores()
	if err != nil {
		handlers.HandleError(c, err)
	}

	c.JSON(http.StatusOK, gin.H{"stores": stores})
}

func GetStoreByID(c *gin.Context) {
	storeStrID := c.Param("id")
	storeID, err := strconv.Atoi(storeStrID)
	if err != nil {
		handlers.HandleError(c, errs.ErrValidationFailed)
		return
	}

	OurStore, err := storeRepository.GetStoreByID(uint(storeID))
	if err != nil {
		handlers.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"store": OurStore})
}

func CreateStore(c *gin.Context) {
	var OurStore models.Store
	if err := c.ShouldBindJSON(&OurStore); err != nil {
		handlers.HandleError(c, errs.ErrValidationFailed)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)

	OurStore.OwnerID = userID

	if err := storeService.CreateStore(OurStore); err != nil {
		handlers.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "store created successfully"})
}

func UpdateStore(c *gin.Context) {
	var OurStore models.Store
	if err := c.ShouldBindJSON(&OurStore); err != nil {
		handlers.HandleError(c, errs.ErrValidationFailed)
		return
	}

	storeStrID := c.Param("storeID")
	storeID, err := strconv.Atoi(storeStrID)
	if err != nil {
		handlers.HandleError(c, errs.ErrValidationFailed)
		return
	}

	store, err := storeRepository.GetStoreByID(uint(storeID))
	if err != nil {
		handlers.HandleError(c, err)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)

	if userID == store.OwnerID {
		handlers.HandleError(c, errs.ErrPermissionDenied)
		return
	}

	OurStore.ID = uint(storeID)
	err = storeRepository.UpdateStore(uint(storeID), &OurStore)
	if err != nil {
		handlers.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "store updated successfully"})
}

func DeleteStore(c *gin.Context) {
	storeStrID := c.Param("storeID")
	storeID, err := strconv.Atoi(storeStrID)
	if err != nil {
		handlers.HandleError(c, errs.ErrValidationFailed)
		return
	}

	store, err := storeRepository.GetStoreByID(uint(storeID))
	if err != nil {
		handlers.HandleError(c, err)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)

	if userID == store.OwnerID {
		handlers.HandleError(c, errs.ErrPermissionDenied)
		return
	}

	err = storeRepository.DeleteStore(uint(storeID))
	if err != nil {
		handlers.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "store deleted successfully"})
}
