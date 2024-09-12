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

// GetAccountsByUserID godoc
// @Summary Get accounts by user ID
// @Description Retrieve a list of accounts for the authenticated user
// @Tags accounts
// @Produce json
// @Success 200 {object} models.Account
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accounts [get]
// @Security ApiKeyAuth
func GetAccountsByUserID(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	accounts, err := repository.GetAccountsByUserID(userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

// GetAccountByID godoc
// @Summary Get account by ID
// @Description Retrieve an account by its ID for the authenticated user
// @Tags accounts
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} models.Account
// @Failure 400 {object} models.ErrorResponse "Invalid Account ID"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 403 {object} models.ErrorResponse "Permission Denied"
// @Failure 404 {object} models.ErrorResponse "Account Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accounts/{id} [get]
// @Security ApiKeyAuth
func GetAccountByID(c *gin.Context) {
	accountStrID := c.Param("id")
	if accountStrID == "" {
		HandleError(c, errs.ErrInvalidAccountID)
		return
	}

	accountID, err := strconv.ParseUint(accountStrID, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrInvalidAccountID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	account, err := repository.GetAccountByID(uint(accountID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrAccountNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	if account.UserID != userID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	c.JSON(http.StatusOK, gin.H{"account": account})
}

// CreateAccount godoc
// @Summary Create a new account
// @Description Create a new account for the authenticated user
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body models.AccountRequest true "Account Data"
// @Success 200 {object} models.DefaultResponse "Account created successfully"
// @Failure 400 {object} models.ErrorResponse "Validation Error"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accounts [post]
// @Security ApiKeyAuth
func CreateAccount(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	var account models.Account
	if err := c.BindJSON(&account); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.ValidateAccount(HandleError, account, c); err != nil {
		return
	}

	account.UserID = userID

	if _, err := repository.GetAccountByNumber(account.AccountNumber); err != nil {
		HandleError(c, errs.ErrAccountNumberUniquenessFailed)
		return
	}

	if err := repository.CreateAccount(account); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account created successfully"})
}

// UpdateAccount godoc
// @Summary Update an existing account
// @Description Update an existing account for the authenticated user
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Param account body models.AccountRequest true "Account Data"
// @Success 200 {object} models.DefaultResponse "Account updated successfully"
// @Failure 400 {object} models.ErrorResponse "Validation Error"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 403 {object} models.ErrorResponse "Permission Denied"
// @Failure 404 {object} models.ErrorResponse "Account Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accounts/{id} [put]
// @Security ApiKeyAuth
func UpdateAccount(c *gin.Context) {
	accountStrID := c.Param("id")
	if accountStrID == "" {
		HandleError(c, errs.ErrInvalidAccountID)
		return
	}

	accountID, err := strconv.ParseUint(accountStrID, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrInvalidAccountID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	var account models.Account
	if err = c.BindJSON(&account); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.ValidateAccount(HandleError, account, c); err != nil {
		return
	}

	if _, err := repository.GetAccountByNumber(account.AccountNumber); err != nil {
		HandleError(c, errs.ErrAccountNumberUniquenessFailed)
		return
	}

	accountData, err := repository.GetAccountByID(uint(accountID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrAccountNotFound)
			return
		}

		HandleError(c, err)
	}

	if accountData.UserID != userID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	account.UserID = userID
	account.ID = uint(accountID)

	err = repository.UpdateAccount(account)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account updated successfully"})
}

// DeleteAccount godoc
// @Summary Delete an account
// @Description Delete an account by its ID for the authenticated user
// @Tags accounts
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} models.DefaultResponse "Account deleted successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid Account ID"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 403 {object} models.ErrorResponse "Permission Denied"
// @Failure 404 {object} models.ErrorResponse "Account Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accounts/{id} [delete]
// @Security ApiKeyAuth
func DeleteAccount(c *gin.Context) {
	accountStrID := c.Param("id")
	if accountStrID == "" {
		HandleError(c, errs.ErrInvalidAccountID)
		return
	}

	accountID, err := strconv.ParseUint(accountStrID, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrInvalidAccountID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	accountData, err := repository.GetAccountByID(uint(accountID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrAccountNotFound)
			return
		}

		HandleError(c, err)
	}

	if accountData.UserID != userID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	err = repository.DeleteAccount(accountData)
	if err != nil {
		HandleError(c, errs.ErrDeleteFailed)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
}

// FillAccountBalance godoc
// @Summary Fills an account
// @Description Fills an account by its account number for the authenticated user
// @Tags accounts
// @Produce json
// @Param account body models.FillAccountRequest true "Account Data"
// @Success 200 {object} models.DefaultResponse "Account filled successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid Account ID"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 403 {object} models.ErrorResponse "Permission Denied"
// @Failure 404 {object} models.ErrorResponse "Account Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /accounts/fill [put]
// @Security ApiKeyAuth
func FillAccountBalance(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}
	var err error

	var account models.Account
	if err = c.BindJSON(&account); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if err = service.ValidateAccount(HandleError, account, c); err != nil {
		return
	}

	accountData, err := repository.GetAccountByNumber(account.AccountNumber)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrAccountNotFound)
			return
		}

		HandleError(c, err)
	}

	if accountData.UserID != userID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	if err = repository.FillAccountBalance(account.AccountNumber, account.Balance); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account filled successfully"})
}
