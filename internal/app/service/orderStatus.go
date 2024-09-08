package service

import (
	"BizMart/internal/app/models"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
	"errors"
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
