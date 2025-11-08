package service

import (
	"loyalty-api/internal/models"
	"loyalty-api/repository"
	"time"
)

type RedemptionService interface {
	Save(redemption models.Redemption) models.Redemption
	DeleteById(id int) error
	FindAll() []models.Redemption
	FindByUserID(id int) []models.Redemption
	FindByReward(id int) []models.Redemption
}
type redemptionService struct {
	service repository.RedemptionRepository
}

// DeleteById implements RedemptionService.
func (r *redemptionService) DeleteById(id int) error {
	if err := r.service.DeleteById(id); err != nil {
		return err
	}
	return nil
}

// FindAll implements RedemptionService.
func (r *redemptionService) FindAll() []models.Redemption {
	return r.service.FindAll()
}

// FindByReward implements RedemptionService.
func (r *redemptionService) FindByReward(id int) []models.Redemption {
	return r.service.FindByReward(id)
}

// FindByUserID implements RedemptionService.
func (r *redemptionService) FindByUserID(id int) []models.Redemption {
	return r.service.FindByUserID(id)
}

// Save implements RedemptionService.
func (r *redemptionService) Save(redemption models.Redemption) models.Redemption {
	redemption.RedeemedAt = time.Now()
	return r.service.Save(redemption)
}

// Constructeur
func NewRedemptionService(repo repository.RedemptionRepository) RedemptionService {
	return &redemptionService{
		service: repo,
	}
}
