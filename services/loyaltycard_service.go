package service

import (
	"loyalty-api/controller/dto"
	"loyalty-api/internal/models"
	"loyalty-api/repository"
	"time"
)

type LoyaltyCardService interface {
	Save(loyaltyCard models.LoyaltyCard) models.LoyaltyCard
	DeleteById(id int) error
	FindAll() []models.LoyaltyCard
	FindAllByUser(id int) []models.Clients
	FindAllByMerchant(id int) []models.LoyaltyCard
	AddPointsCard(loyalty models.LoyaltyCard, points int) (models.LoyaltyCard, error)
	SoldePointsClient(loyalty models.LoyaltyCard) (int, error)
	RetraitPointsClient(loyalty models.LoyaltyCard, points int) (models.LoyaltyCard, error)
	SoldeMerchant(merchantId int) int
	SoldePointsClientByAllMerchants(clientId int) (int, error)
	FindAllMerchantByClient(clientId int) []dto.MerchantDto
}
type loyaltyCardService struct {
	service repository.LoyaltyRepository
}

// FindAllMerchantByClient implements LoyaltyCardService.
func (l *loyaltyCardService) FindAllMerchantByClient(clientId int) []dto.MerchantDto {
	return l.service.FindAllMerchantByClient(clientId)
}

// SoldePointsClientByAllMerchants implements LoyaltyCardService.
func (l *loyaltyCardService) SoldePointsClientByAllMerchants(clientId int) (int, error) {
	return l.service.SoldePointsClientByAllMerchants(clientId)
}

// SoldeMerchant implements LoyaltyCardService.
func (l *loyaltyCardService) SoldeMerchant(merchantId int) int {
	return l.service.SoldeMerchant(merchantId)
}

// RetraitPointsClient implements LoyaltyCardService.
func (l *loyaltyCardService) RetraitPointsClient(loyalty models.LoyaltyCard, points int) (models.LoyaltyCard, error) {
	loyalty, err := l.service.RetraitPointsClient(loyalty, points)
	if err != nil {
		return loyalty, err
	}
	return loyalty, nil
}

// SoldePointsClient implements LoyaltyCardService.
func (l *loyaltyCardService) SoldePointsClient(loyalty models.LoyaltyCard) (int, error) {
	points, err := l.service.SoldePointsClient(loyalty)
	if err != nil {
		return points, err
	}
	return points, nil
}

// AddPointsCard implements LoyaltyCardService.
func (l *loyaltyCardService) AddPointsCard(loyalty models.LoyaltyCard, points int) (models.LoyaltyCard, error) {
	if loyalty, err := l.service.AddPoints(loyalty, points); err != nil {
		return loyalty, err
	}
	return loyalty, nil
}

// DeleteById implements LoyaltyCardService.
func (l *loyaltyCardService) DeleteById(id int) error {
	if err := l.service.DeleteById(id); err != nil {
		return err
	}
	return nil
}

// FindAll implements LoyaltyCardService.
func (l *loyaltyCardService) FindAll() []models.LoyaltyCard {
	return l.service.FindAll()
}

// FindAllByMerchant implements LoyaltyCardService.
func (l *loyaltyCardService) FindAllByMerchant(id int) []models.LoyaltyCard {
	return l.service.FindAllByMerchant(id)
}

// FindAllByUser implements LoyaltyCardService.
func (l *loyaltyCardService) FindAllByUser(id int) []models.Clients {
	return l.service.FindAllByUser(id)
}

// Save implements LoyaltyCardService.
func (l *loyaltyCardService) Save(loyaltyCard models.LoyaltyCard) models.LoyaltyCard {
	loyaltyCard.LastUpdated = time.Now()
	return l.service.Save(loyaltyCard)
}

// constructeur
func NewLoyaltyCardService(repo repository.LoyaltyRepository) LoyaltyCardService {
	return &loyaltyCardService{
		service: repo,
	}
}
