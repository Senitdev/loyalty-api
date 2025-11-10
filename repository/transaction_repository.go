package repository

import (
	"loyalty-api/internal/models"
	"strings"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(transaction models.Transaction) models.Transaction
	DeleteById(id int) error
	FindAll() []models.Transaction
	FindbyLoyalCard(id int) []models.Transaction
	FindbyMerchant(merchant_id int, startDate, endDate string) ([]models.Transaction, error)
}
type transactionRepository struct {
	DB *gorm.DB
}

// FndbyMerchant implements TransactionRepository.
func (t *transactionRepository) FindbyMerchant(merchant_id int, startDate, endDate string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if len(startDate) > 10 {
		startDate = startDate[:10]
	}
	if len(endDate) > 10 {
		endDate = endDate[:10]
	}
	startDate = strings.TrimSpace(startDate)
	endDate = strings.TrimSpace(endDate)
	query := t.DB.Where("merchant_id = ?", merchant_id)
	if startDate != "" && endDate != "" {
		query = query.Where("DATE(created_at) BETWEEN ? AND ?", startDate, endDate)

	} else {
		// Sinon on limite à 4 résultats
		query = query.Limit(4)
	}
	result := query.Order("created_at DESC").Find(&transactions)
	if result.Error != nil {
		return transactions, result.Error
	}
	return transactions, nil
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
	t.DB.Find(&transaction)
	return transaction
}

// FindbyLoyalCard implements TransactionRepository.
func (t *transactionRepository) FindbyLoyalCard(id int) []models.Transaction {
	var transaction []models.Transaction
	t.DB.Where("loyalty_card_id", id).Find(&transaction)
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
