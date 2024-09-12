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

// GetAddressesByUserID получает список адресов пользователя
// @Summary Get addresses by user ID
// @Description Get all addresses for the authenticated user
// @Tags addresses
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Address
// @Failure 401 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /address [get]
func GetAddressesByUserID(c *gin.Context) {
	// Получаем ID пользователя
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	addresses, err := repository.GetMyAddresses(userID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrAddressNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"addresses": addresses})
}

// GetAddressByID получает адрес по его ID
// @Summary Get address by ID
// @Description Get address details by address ID for the authenticated user
// @Tags addresses
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Address ID"
// @Success 200 {object} models.Address
// @Failure 401 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /address/{id} [get]
func GetAddressByID(c *gin.Context) {
	addressStrID := c.Param("id")
	addressID, err := strconv.ParseUint(addressStrID, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	address, err := repository.GetAddressByID(uint(addressID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrAddressNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	if address.UserID != userID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	c.JSON(http.StatusOK, gin.H{"address": address})
}

// CreateAddress создает новый адрес
// @Summary Create a new address
// @Description Create a new address for the authenticated user
// @Tags addresses
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param address body models.AddressRequest true "Address data"
// @Success 200 {object} models.DefaultResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /address [post]
func CreateAddress(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	var address models.Address
	if err := c.BindJSON(&address); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if _, err := repository.GetAddressByNameAndUserID(address.AddressName, userID); err == nil {
		HandleError(c, errs.ErrAddressNameUniquenessFailed)
		return
	}

	if err := service.ValidateAddress(HandleError, address, c); err != nil {
		return
	}

	address.UserID = userID

	err := repository.CreateAddress(&address)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "address created successfully"})
}

// UpdateAddress обновляет информацию об адресе
// @Summary Update an existing address
// @Description Update address details by address ID for the authenticated user
// @Tags addresses
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path string true "Address ID"
// @Param address body models.AddressRequest true "Updated address data"
// @Success 200 {object} models.AddressRequest
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /address/{id} [put]
func UpdateAddress(c *gin.Context) {
	addressStrID := c.Param("id")
	addressID, err := strconv.ParseUint(addressStrID, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrInvalidAddressID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
	}

	var address models.Address
	if err = c.BindJSON(&address); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if _, err := repository.GetAddressByNameAndUserID(address.AddressName, userID); err == nil {
		HandleError(c, errs.ErrAddressNameUniquenessFailed)
		return
	}

	if err = service.ValidateAddress(HandleError, address, c); err != nil {
		return
	}

	addressData, err := repository.GetAddressByID(uint(addressID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrAddressNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	if addressData.UserID != userID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	address.ID = addressData.ID
	address.UserID = userID

	err = repository.UpdateAddress(&address)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrAddressNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "address updated successfully"})
}

// DeleteAddress удаляет адрес
// @Summary Delete an address
// @Description Delete address by ID for the authenticated user
// @Tags addresses
// @Security ApiKeyAuth
// @Param id path string true "Address ID"
// @Success 200 {object} models.DefaultResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /address/{id} [delete]
func DeleteAddress(c *gin.Context) {
	addressStrID := c.Param("id")
	addressID, err := strconv.ParseUint(addressStrID, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrInvalidAddressID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	address, err := repository.GetAddressByID(uint(addressID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrAddressNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	if address.UserID != userID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	err = repository.DeleteAddress(address.ID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrAddressNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "address deleted successfully"})
}
