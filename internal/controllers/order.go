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
// @Success 200 {object} models.Order "orders"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
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
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// GetOrderByID godoc
// @Summary Get order by ID
// @Description Retrieves a specific order by its ID if it belongs to the authenticated user.
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order "order"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Order not found"
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

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	if order.UserID != userID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Allows the authenticated user to create a new order.
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body models.OrderRequestJsonBind true "Order Data"
// @Success 201 {object} models.DefaultResponse "Order created successfully"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 400 {object} models.ErrorResponse "Validation failed"
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

	orderRequest.UserID = userID
	orderRequest.StatusID = 1

	if err := service.ValidateOrder(HandleError, orderRequest, c); err != nil {
		return
	}

	if err := service.CreateOrder(orderRequest); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "order created successfully"})
}

// UpdateOrder godoc
// @Summary Update an order
// @Description Allows the authenticated user to update an existing order.
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Param order body models.OrderRequestJsonBind true "Updated Order Data"
// @Success 200 {object} models.DefaultResponse "Order updated successfully"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Order not found"
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

	order, err := repository.GetOrderByID(uint(orderId))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrOrderNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	if _, err = repository.GetPaymentByOrderID(order.ID); err == nil {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	if order.UserID != userID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	orderDetails, err := repository.GetOrderDetailsByID(order.OrderDetailsID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrOrderNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	orderRequest.ProductID = orderDetails.ProductID

	if err := service.UpdateOrder(uint(orderId), orderRequest); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order updated successfully"})
}

// DeleteOrder godoc
// @Summary Delete an order
// @Description Allows the authenticated user to delete an order.
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} models.DefaultResponse "Order deleted successfully"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Order not found"
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

	if err := service.DeleteOrder(userID, uint(orderID)); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order deleted successfully"})
}
