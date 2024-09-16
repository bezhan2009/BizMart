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

// GetAllUserOrders godoc
// @Summary Get all orders for a user
// @Description Retrieves all orders associated with the authenticated user.
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {object} models.DefaultResponse "orders retrieved successfully"
// @Failure 401 {object} models.ErrorResponse "unauthorized access"
// @Failure 500 {object} models.ErrorResponse "internal server error"
// @Router /orders [get]
func GetAllUserOrders(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	orders, err := repository.GetAllOrderByUserID(userID)
	if err != nil {
		HandleError(c, err)
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// GetOrderByID godoc
// @Summary Get order by ID
// @Description Retrieves a specific order by its ID.
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path uint true "Order ID"
// @Success 200 {object} models.DefaultResponse "order retrieved successfully"
// @Failure 400 {object} models.ErrorResponse "invalid order ID"
// @Failure 404 {object} models.ErrorResponse "order not found"
// @Failure 500 {object} models.ErrorResponse "internal server error"
// @Router /orders/{id} [get]
func GetOrderByID(c *gin.Context) {
	orderIdStr := c.Param("id")
	if orderIdStr == "" {
		HandleError(c, errs.ErrInvalidOrderID)
		return
	}

	orderID, err := strconv.ParseUint(orderIdStr, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrInvalidOrderID)
		return
	}

	if orderID == 0 {
		HandleError(c, errs.ErrInvalidOrderID)
		return
	}

	order, err := repository.GetOrderByID(uint(orderID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrOrderNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Creates a new order with the provided details.
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body models.OrderRequest true "Order Request"
// @Success 201 {object} models.DefaultResponse "order created successfully"
// @Failure 400 {object} models.ErrorResponse "invalid request data"
// @Failure 401 {object} models.ErrorResponse "unauthorized access"
// @Failure 500 {object} models.ErrorResponse "internal server error"
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	var orderRequest models.OrderRequestJsonBind
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	product, err := repository.GetProductByID(orderRequest.ProductID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrProductNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	if err = service.ValidateOrder(HandleError, orderRequest, product, c); err != nil {
		return
	}

	if err := service.CreateOrder(orderRequest); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "order created successfully"})
}

// UpdateOrder godoc
// @Summary Update an existing order
// @Description Updates an order with the provided details.
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path uint true "Order ID"
// @Param order body models.OrderRequest true "Order Request"
// @Success 200 {object} models.DefaultResponse "order updated successfully"
// @Failure 400 {object} models.ErrorResponse "invalid request data"
// @Failure 401 {object} models.ErrorResponse "unauthorized access"
// @Failure 404 {object} models.ErrorResponse "order not found"
// @Failure 500 {object} models.ErrorResponse "internal server error"
// @Router /orders/{id} [put]
func UpdateOrder(c *gin.Context) {
	orderIdStr := c.Param("id")
	orderId, err := strconv.ParseUint(orderIdStr, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrInvalidOrderID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	var orderRequest models.OrderRequestJsonBind
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	product, err := repository.GetProductByID(orderRequest.ProductID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrProductNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	if err = service.ValidateOrder(HandleError, orderRequest, product, c); err != nil {
		return
	}

	order, err := repository.GetOrderByID(uint(orderId))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrOrderNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	order.ID = uint(orderId)
	if order.UserID != userID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	if err := service.UpdateOrder(uint(orderId), orderRequest); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "order updated successfully"})
}

// DeleteOrder godoc
// @Summary Delete an order
// @Description Deletes an order by its ID.
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path uint true "Order ID"
// @Success 200 {object} models.DefaultResponse "order deleted successfully"
// @Failure 400 {object} models.ErrorResponse "invalid order ID"
// @Failure 404 {object} models.ErrorResponse "order not found"
// @Failure 500 {object} models.ErrorResponse "internal server error"
// @Router /orders/{id} [delete]
func DeleteOrder(c *gin.Context) {
	orderIdStr := c.Param("id")
	if orderIdStr == "" {
		HandleError(c, errs.ErrInvalidOrderID)
		return
	}

	orderID, err := strconv.ParseUint(orderIdStr, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrInvalidOrderID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	order, err := repository.GetOrderByID(uint(orderID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrOrderNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	order.ID = uint(orderID)
	if order.UserID != userID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	if err := repository.DeleteOrder(uint(orderID)); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order deleted successfully"})
}
