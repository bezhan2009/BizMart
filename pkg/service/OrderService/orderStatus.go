package OrderService

import (
	"BizMart/errs"
	"BizMart/models"
	"BizMart/pkg/repository/orderRepository"
)

func CreateOrderStatus(orderStatus models.OrderStatus) (orderStatusID uint, err error) {
	var orderStat models.OrderStatus

	orderStat, err = orderRepository.GetOrderStatusByName(orderStatus.StatusName)
	if orderStat.ID != 0 {
		return 0, errs.ErrOrderStatusNameUniquenessFailed
	}

	orderStatID, err := orderRepository.CreateOrderStatus(orderStatus)
	if err != nil {
		return 0, errs.ErrValidationFailed
	}

	return orderStatID, nil
}
