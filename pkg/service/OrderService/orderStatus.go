package OrderService

import (
	"BizMart/errs"
	"BizMart/models"
	"BizMart/pkg/repository/orderRepository"
	"errors"
)

func CreateOrderStatus(orderStatus models.OrderStatus) (orderStatusID uint, err error) {
	var orderStat models.OrderStatus

	orderStat, err = orderRepository.GetOrderStatusByName(orderStatus.StatusName)
	if err == nil {
		return 0, errs.ErrOrderStatusNameUniquenessFailed
	}
	if orderStat.ID != 0 {
		return 0, errs.ErrOrderStatusNameUniquenessFailed
	}

	orderStatID, err := orderRepository.CreateOrderStatus(orderStatus)
	if err != nil {
		return 0, errs.ErrValidationFailed
	}

	return orderStatID, nil
}

func UpdateOrderStatus(ordStatID uint, orderStatus models.OrderStatus) (orderStatusID uint, err error) {
	_, err = orderRepository.GetOrderStatusByID(ordStatID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return 0, errs.ErrOrderStatusNotFound
		}

		return 0, err
	}

	orderStatusID, err = orderRepository.UpdateOrderStatus(ordStatID, orderStatus)
	if err != nil {
		return 0, errs.ErrValidationFailed
	}

	return orderStatusID, nil
}
