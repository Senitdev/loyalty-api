package controller

import (
	"loyalty-api/internal/models"
	service "loyalty-api/services"

	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	Save(ctx *gin.Context) models.Transaction
	DeleteById(id int) error
	FindAll() []models.Transaction
	FindbyLoyalCard(id int) []models.Transaction
	FindbyMerchant(merchant_id int, startDate, endDate string) ([]models.Transaction, error)
}
type transactionController struct {
	controller service.TransactionService
}

// FindbyMerchant implements TransactionController.
func (t *transactionController) FindbyMerchant(id int, startDate, endDate string) ([]models.Transaction, error) {
	return t.controller.FindbyMerchant(id, startDate, endDate)
}

// DeleteById implements TransactionController.
func (t *transactionController) DeleteById(id int) error {
	if err := t.controller.DeleteById(id); err != nil {
		return err
	}
	return nil
}

// FindAll implements TransactionController.
func (t *transactionController) FindAll() []models.Transaction {
	return t.controller.FindAll()
}

// FindbyLoyalCard implements TransactionController.
func (t *transactionController) FindbyLoyalCard(id int) []models.Transaction {
	return t.controller.FindbyLoyalCard(id)
}

// Save implements TransactionController.
func (t *transactionController) Save(ctx *gin.Context) models.Transaction {
	var transaction models.Transaction
	ctx.BindJSON(&transaction)
	t.controller.Save(transaction)
	return transaction
}

// Constructeur
func NewTransactionController(service service.TransactionService) TransactionController {
	return &transactionController{
		controller: service,
	}
}
