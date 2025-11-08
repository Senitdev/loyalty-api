package repository

import (
	"loyalty-api/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(transaction models.Transaction) models.Transaction
	DeleteById(id int) error
	FindAll() []models.Transaction
	FindbyLoyalCard(id int) []models.Transaction
	FindbyMerchant(id int) []models.Transaction
}
type transactionRepository struct {
	DB *gorm.DB
}

// FndbyMerchant implements TransactionRepository.
func (t *transactionRepository) FindbyMerchant(id int) []models.Transaction {
	var transaction []models.Transaction
	t.DB.Where("merchant_id", id).Find(&models.Transaction{})
	return transaction
}

// DeleteById implements TransactionRepository.
func (t *transactionRepository) DeleteById(id int) error {
	if result := t.DB.Delete(&models.Transaction{}, id).Error; result != nil {
		return result
	}
	return nil
}

// FindAll implements TransactionRepository.
func (t *transactionRepository) FindAll() []models.Transaction {
	var transaction []models.Transaction
	t.DB.Find(&models.Transaction{})
	return transaction
}

// FindbyLoyalCard implements TransactionRepository.
func (t *transactionRepository) FindbyLoyalCard(id int) []models.Transaction {
	var transaction []models.Transaction
	t.DB.Where("loyalty_card_id", id).Find(&models.Transaction{})
	return transaction
}

// Save implements TransactionRepository.
func (t *transactionRepository) Save(transaction models.Transaction) models.Transaction {
	t.DB.Save(&transaction)
	return transaction
}

// constructeur
func NewTransctionRepository(repo *gorm.DB) TransactionRepository {
	return &transactionRepository{
		DB: repo,
	}
}
