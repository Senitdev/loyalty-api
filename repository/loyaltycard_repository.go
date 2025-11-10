package repository

import (
	"errors"
	"loyalty-api/internal/models"
	"time"

	"gorm.io/gorm"
)

type LoyaltyRepository interface {
	Save(loyaltyCard models.LoyaltyCard) models.LoyaltyCard
	DeleteById(id int) error
	FindAll() []models.LoyaltyCard
	FindAllByUser(id int) []models.LoyaltyCard
	FindAllByMerchant(id int) []models.LoyaltyCard
	AddPoints(loyaly models.LoyaltyCard, points int) (models.LoyaltyCard, error)
	SoldePointsClient(loyalty models.LoyaltyCard) (int, error)
	RetraitPointsClient(loyalty models.LoyaltyCard, points int) (models.LoyaltyCard, error)
	SoldeMerchant(merchantId int) int
}
type loyaltyCardRepository struct {
	BD *gorm.DB
}

// SoldeMerchant implements LoyaltyRepository.
func (l *loyaltyCardRepository) SoldeMerchant(merchantId int) int {
	var solde int64
	err := l.BD.Model(&models.LoyaltyCard{}).
		Where("merchant_id = ?", merchantId).
		Select("COALESCE(SUM(points), 0)").Scan(&solde).Error
	if err != nil {
		return 0
	}
	return int(solde)
}

// RetraitPointsClient implements LoyaltyRepository.
func (l *loyaltyCardRepository) RetraitPointsClient(loyalty models.LoyaltyCard, points int) (models.LoyaltyCard, error) {
	var transaction models.Transaction
	result := l.BD.Where("clients_id = ? AND merchant_id = ?", loyalty.ClientsID, loyalty.MerchantID).First(&loyalty)
	if result.Error != nil {
		return loyalty, result.Error
	}
	if loyalty.Points >= points {
		loyalty.Points -= points
		l.BD.Save(loyalty)
		//On doit enregistrer l historique
		transaction.CreatedAt = time.Now()
		transaction.Description = "Consommation"
		transaction.LoyaltyCardID = loyalty.ID
		transaction.MerchantId = loyalty.MerchantID
		transaction.Points = points
		transaction.Type = "spen"
		l.BD.Save(&transaction)
	} else {
		return loyalty, result.Error
	}
	return loyalty, nil
}

// SoldePointsClient implements LoyaltyRepository.
func (l *loyaltyCardRepository) SoldePointsClient(loyalty models.LoyaltyCard) (int, error) {
	// Recherche d’une carte existante pour ce client et ce marchand
	result := l.BD.Where("clients_id = ? AND merchant_id = ?", loyalty.ClientsID, loyalty.MerchantID).First(&loyalty)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, result.Error
		}
		return 0, result.Error
	}
	return loyalty.Points, nil
}

// AddPoints implements LoyaltyRepository.
func (l *loyaltyCardRepository) AddPoints(loyalty models.LoyaltyCard, points int) (models.LoyaltyCard, error) {
	var existing models.LoyaltyCard

	// Recherche d’une carte existante pour ce client et ce marchand
	result := l.BD.Where("clients_id = ? AND merchant_id = ?", loyalty.ClientsID, loyalty.MerchantID).First(&existing)
	if result.Error != nil {
		// Cas 1️⃣ : aucune carte trouvée → on en crée une nouvelle
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			loyalty.Points = points
			loyalty.LastUpdated = time.Now()

			if err := l.BD.Create(&loyalty).Error; err != nil {
				return models.LoyaltyCard{}, err
			}
			return loyalty, nil
		}

		// Cas 2️⃣ : autre erreur SQL → on la renvoie
		return models.LoyaltyCard{}, result.Error
	}

	// Cas 3️⃣ : carte trouvée → on ajoute les points
	existing.Points += points
	existing.LastUpdated = time.Now()

	if err := l.BD.Save(&existing).Error; err != nil {
		return models.LoyaltyCard{}, err
	}
	var transaction models.Transaction
	//On doit enregistrer l historique
	transaction.CreatedAt = time.Now()
	transaction.Description = "Earn"
	transaction.LoyaltyCardID = existing.ID
	transaction.MerchantId = existing.MerchantID
	transaction.Points = points
	transaction.Type = "earn"
	l.BD.Save(&transaction)
	return existing, nil
}

// DeleteById implements LoyaltyRepository.
func (l *loyaltyCardRepository) DeleteById(id int) error {
	if result := l.BD.Delete(&models.LoyaltyCard{}, id).Error; result != nil {
		return result
	}
	return nil
}

// FindAll implements LoyaltyRepository.
func (l *loyaltyCardRepository) FindAll() []models.LoyaltyCard {
	var loyaltyCard []models.LoyaltyCard
	result := l.BD.Find(&models.LoyaltyCard{})
	if result.Error != nil {
		return loyaltyCard
	}
	return loyaltyCard
}

// FindAllByMerchant implements LoyaltyRepository.
func (l *loyaltyCardRepository) FindAllByMerchant(id int) []models.LoyaltyCard {
	var loyaltyCard []models.LoyaltyCard
	result := l.BD.Where("merchant_id", id).Find(&models.LoyaltyCard{})
	if result.Error != nil {
		return loyaltyCard
	}
	return loyaltyCard
}

// FindAllByUser implements LoyaltyRepository.
func (l *loyaltyCardRepository) FindAllByUser(id int) []models.LoyaltyCard {
	var loyaltyCard []models.LoyaltyCard
	result := l.BD.Where("user_id", id).Find(&models.LoyaltyCard{})
	if result.Error != nil {
		return loyaltyCard
	}
	return loyaltyCard
}

// Save implements LoyaltyRepository.
func (l *loyaltyCardRepository) Save(loyaltyCard models.LoyaltyCard) models.LoyaltyCard {
	l.BD.Save(&loyaltyCard)
	return loyaltyCard
}

// constructeur
func NewLoyaltyCardRepository(repo *gorm.DB) LoyaltyRepository {
	return &loyaltyCardRepository{
		BD: repo,
	}
}
