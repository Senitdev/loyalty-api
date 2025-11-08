package service

import (
	"loyalty-api/internal/models"
	"loyalty-api/repository"
	"time"
)

type TransactionService interface {
	Save(transaction models.Transaction) models.Transaction
	DeleteById(id int) error
	FindAll() []models.Transaction
	FindbyLoyalCard(id int) []models.Transaction
	FindbyMerchant(id int) []models.Transaction
}
type transactionService struct {
	service repository.TransactionRepository
}

// FindbyMerchant implements TransactionService.
func (t *transactionService) FindbyMerchant(id int) []models.Transaction {
	return t.service.FindbyMerchant(id)
}

// DeleteById implements TransactionService.
func (t *transactionService) DeleteById(id int) error {
	if err := t.service.DeleteById(id); err != nil {
		return err
	}
	return nil
}

// FindAll implements TransactionService.
func (t *transactionService) FindAll() []models.Transaction {
	return t.service.FindAll()
}

// FindbyLoyalCard implements TransactionService.
func (t *transactionService) FindbyLoyalCard(id int) []models.Transaction {
	return t.service.FindbyLoyalCard(id)
}

// Save implements TransactionService.
func (t *transactionService) Save(transaction models.Transaction) models.Transaction {
	transaction.CreatedAt = time.Now()
	return t.service.Save(transaction)
}

// Constructeur
func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{
		service: repo,
	}
}
