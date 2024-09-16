package service

import (
	"BizMart/internal/app/models"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
	"errors"
	"github.com/gin-gonic/gin"
)

func CreateOrderStatus(orderStatus models.OrderStatus) (orderStatusID uint, err error) {
	var orderStat models.OrderStatus

	orderStat, err = repository.GetOrderStatusByName(orderStatus.StatusName)
	if err == nil {
		return 0, errs.ErrOrderStatusNameUniquenessFailed
	}
	if orderStat.ID != 0 {
		return 0, errs.ErrOrderStatusNameUniquenessFailed
	}

	orderStatID, err := repository.CreateOrderStatus(orderStatus)
	if err != nil {
		return 0, errs.ErrValidationFailed
	}

	return orderStatID, nil
}

func UpdateOrderStatus(ordStatID uint, orderStatus models.OrderStatus) (orderStatusID uint, err error) {
	_, err = repository.GetOrderStatusByID(ordStatID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return 0, errs.ErrOrderStatusNotFound
		}

		return 0, err
	}

	orderStatusID, err = repository.UpdateOrderStatus(ordStatID, orderStatus)
	if err != nil {
		return 0, errs.ErrValidationFailed
	}

	return orderStatusID, nil
}

func CreateOrder(orderRequest models.OrderRequestJsonBind) (err error) {
	var order models.Order
	var orderDetails models.OrderDetails

	order.UserID = orderRequest.UserID
	order.StatusID = 1
	orderDetails.ProductID = orderRequest.ProductID
	orderDetails.AddressID = orderRequest.AddressID
	orderDetails.Quantity = orderRequest.Quantity

	product, err := repository.GetProductByID(orderDetails.ProductID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrProductNotFound
		}

		return err
	}

	orderDetails.Price = product.Price * float64(orderRequest.Quantity)

	_, err = repository.GetAddressByID(orderDetails.AddressID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrProductNotFound
		}

		return err
	}

	if err = repository.CreateOrder(order, orderDetails); err != nil {
		return err
	}

	return nil
}

func UpdateOrder(orderID uint, orderRequest models.OrderRequestJsonBind) (err error) {
	var order models.Order
	var orderDetails models.OrderDetails

	order.ID = orderID

	product, err := repository.GetProductByID(orderRequest.ProductID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrProductNotFound
		}

		return err
	}

	order, err = repository.GetOrderByID(orderID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrOrderNotFound
		}

		return err
	}

	orderDetails, err = repository.GetOrderDetailsByID(order.OrderDetailsID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrOrderNotFound
		}

		return err
	}

	orderDetails.Price = product.Price * float64(orderRequest.Quantity)
	orderDetails.AddressID = orderRequest.AddressID
	order.StatusID = orderRequest.StatusID

	if err = repository.UpdateOrder(order, orderDetails); err != nil {
		return err
	}

	return nil
}

func ValidateOrder(HandleError func(ctx *gin.Context, err error), orderData models.OrderRequestJsonBind, productData models.Product, c *gin.Context) error {
	if orderData.Quantity > productData.Amount {
		HandleError(c, errs.ErrInvalidQuantity)
		return errs.ErrInvalidQuantity
	}

	if _, err := repository.GetAddressByID(orderData.AddressID); err != nil {
		HandleError(c, errs.ErrAddressNotFound)
		return err
	}
}
