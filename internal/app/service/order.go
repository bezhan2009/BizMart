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

	order.StatusID = 1

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

	if orderRequest.Quantity != 0 {
		// Проверка на отрицательное значение
		if product.Amount < 0 || orderRequest.Quantity > product.Amount {
			return errs.ErrNotEnoughProductInStock
		}

		newAmount := product.Amount + orderDetails.Quantity - orderRequest.Quantity

		product.Amount = newAmount

		if err = repository.UpdateProduct(&product); err != nil {
			return err
		}
	}

	orderDetails.Price = product.Price * float64(orderRequest.Quantity)
	orderDetails.AddressID = orderRequest.AddressID
	order.StatusID = orderRequest.StatusID

	if err = repository.UpdateOrder(order, orderDetails); err != nil {
		return err
	}

	return nil
}

func DeleteOrder(userID, orderID uint) (err error) {
	order, err := repository.GetOrderByID(orderID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrOrderNotFound
		}

		return err
	}

	if order.UserID != userID {
		return errs.ErrPermissionDenied
	}

	if order.StatusID == 3 || order.StatusID == 4 {
		orderDetails, err := repository.GetOrderDetailsByID(order.OrderDetailsID)
		if err != nil {
			if errors.Is(err, errs.ErrRecordNotFound) {
				return errs.ErrOrderNotFound
			}

			return err
		}

		product, err := repository.GetProductByID(orderDetails.ProductID)
		if err != nil {
			if errors.Is(err, errs.ErrRecordNotFound) {
				return errs.ErrProductNotFound
			}

			return err
		}

		product.Amount += orderDetails.Quantity
	}

	if err = repository.DeleteOrder(orderID); err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrOrderNotFound
		}

		return err
	}

	return nil
}

func ValidateOrder(HandleError func(ctx *gin.Context, err error), orderData models.OrderRequestJsonBind, c *gin.Context) error {
	var product models.Product
	var address *models.Address
	var err error

	if address, err = repository.GetAddressByID(orderData.AddressID); err != nil {
		HandleError(c, errs.ErrAddressNotFound)
		return errs.ErrAddressNotFound
	}

	if address.UserID != orderData.UserID {
		HandleError(c, errs.ErrAddressNotFound)
		return errs.ErrAddressNotFound
	}

	if product, err = repository.GetProductByID(orderData.ProductID); err != nil {
		HandleError(c, errs.ErrProductNotFound)
		return errs.ErrProductNotFound
	}

	if _, err := repository.GetOrderStatusByID(orderData.StatusID); err != nil {
		HandleError(c, errs.ErrOrderStatusNotFound)
		return errs.ErrOrderStatusNotFound
	}

	if orderData.Quantity > product.Amount || orderData.Quantity > 1000 {
		HandleError(c, errs.ErrInvalidQuantity)
		return errs.ErrInvalidQuantity
	}

	return nil
}
