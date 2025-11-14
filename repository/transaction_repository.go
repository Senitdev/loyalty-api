package repository

import (
	"fmt"
	"loyalty-api/controller/dto"
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
	FindByClient(client_id int, startDate, endDate string) ([]dto.TransactionDTO, error)
}
type transactionRepository struct {
	DB *gorm.DB
}

func (t *transactionRepository) FindByClient(client_id int, startDate string, endDate string) ([]dto.TransactionDTO, error) {
	var transactions []models.Transaction
	var transactionDTOs []dto.TransactionDTO

	// ------------------------------
	// 1. Nettoyage des dates
	// ------------------------------
	if len(startDate) > 10 {
		startDate = startDate[:10]
	}
	if len(endDate) > 10 {
		endDate = endDate[:10]
	}

	startDate = strings.TrimSpace(startDate)
	endDate = strings.TrimSpace(endDate)
	fmt.Println("startDate:", startDate, "endDate:", endDate)
	// ------------------------------
	// 2. Requête de base
	// ------------------------------
	query := t.DB.Where("client_id = ?", client_id)

	if startDate != "20250101" && endDate != "" && startDate != "" {
		query = query.Where("DATE(created_at) BETWEEN ? AND ?", startDate, endDate)
	} else {
		query = query.Limit(4)
	}

	result := query.Order("created_at DESC").Find(&transactions)
	if result.Error != nil {
		return transactionDTOs, result.Error
	}

	// --------------------------------------------
	// 3. Charger tous les merchantId utilisés
	// --------------------------------------------
	merchantIds := make([]int, 0)
	for _, tr := range transactions {
		merchantIds = append(merchantIds, tr.MerchantId)
	}

	// --------------------------------------------
	// 4. Charger tous les marchands d'un coup
	// --------------------------------------------
	var merchants []models.Merchant
	t.DB.Where("id IN ?", merchantIds).Find(&merchants)

	// --------------------------------------------
	// 5. Construire une map ID → merchant
	// --------------------------------------------
	merchantMap := make(map[int]models.Merchant)
	for _, m := range merchants {
		merchantMap[m.ID] = m
	}

	// --------------------------------------------
	// 6. Construire les DTO
	// --------------------------------------------
	for _, tr := range transactions {
		mName := ""
		if m, ok := merchantMap[tr.MerchantId]; ok {
			mName = m.Name
		}

		dto := dto.TransactionDTO{
			ID:         tr.ID,
			Type:       tr.Type,
			Points:     tr.Points,
			CreatedAt:  tr.CreatedAt,
			MerchantId: tr.MerchantId,
			ClientId:   tr.ClientId,
			Merchant:   mName,
		}

		transactionDTOs = append(transactionDTOs, dto)
	}

	return transactionDTOs, nil
}

// DeleteById implements TransactionRepository.
func (t *transactionRepository) DeleteById(id int) error {
	if result := t.DB.Delete(&models.Transaction{}, id).Error; result != nil {
		return result
	}
	return nil
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
