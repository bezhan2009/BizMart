package service

import (
	"BizMart/internal/app/models"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
	"errors"
	"github.com/gin-gonic/gin"
)

func ValidatePayment(HandleError func(ctx *gin.Context, err error), paymentData *models.Payment, c *gin.Context) error {
	account, err := repository.GetAccountByID(paymentData.AccountID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrAccountNotFound)
			return errs.ErrAccountNotFound
		}

		HandleError(c, err)
		return err
	}

	if account.UserID != paymentData.UserID {
		HandleError(c, errs.ErrAccountNotFound)
		return errs.ErrAccountNotFound
	}

	if paymentData.OrderID == 0 {
		HandleError(c, errs.ErrInvalidOrderID)
		return errs.ErrInvalidOrderID
	}

	order, err := repository.GetOrderByID(paymentData.OrderID)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrOrderNotFound)
			return errs.ErrOrderNotFound
		}

		HandleError(c, err)
		return err
	}

	if order.UserID != paymentData.UserID {
		HandleError(c, errs.ErrOrderNotFound)
		return errs.ErrOrderNotFound
	}

	if order.OrderDetails.Price != paymentData.Price {
		paymentData.Price = order.OrderDetails.Price
	}

	if order.OrderDetails.Quantity != paymentData.Amount {
		paymentData.Amount = order.OrderDetails.Quantity
	}

	return nil
}

func CreatePayment(payment models.Payment) error {
	order, err := repository.GetOrderByID(payment.OrderID)
	if err != nil {
		return err
	}

	if order.StatusID == 3 || order.StatusID == 4 {
		return errs.ErrOrderAlreadyPayed
	}

	var accountStore models.Account

	account, err := repository.GetAccountByID(payment.AccountID)
	if err != nil {
		return err
	}

	product, err := repository.GetProductByID(order.OrderDetails.ProductID)
	if err != nil {
		return err
	}

	if account.Balance > order.OrderDetails.Price {
		account.Balance -= order.OrderDetails.Price
		store, err := repository.GetStoreByID(product.StoreID)
		if err != nil {
			return err
		}

		accountStores, err := repository.GetAccountsByUserID(store.OwnerID)
		if err != nil {
			return err
		}

		accountStore = accountStores[0]
		accountStore.Balance += order.OrderDetails.Price
	} else {
		return errs.ErrInsufficientFunds
	}

	if err = repository.UpdateAccount(account); err != nil {
		return err
	}
	if err = repository.UpdateAccount(accountStore); err != nil {
		return err
	}

	order.StatusID = 3
	if err = repository.UpdateOrder(order, order.OrderDetails); err != nil {
		return err
	}

	if err = repository.CreatePayment(payment); err != nil {
		return err
	}

	return nil
}
