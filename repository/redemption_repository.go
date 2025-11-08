package repository

import (
	"loyalty-api/internal/models"

	"gorm.io/gorm"
)

type RedemptionRepository interface {
	Save(redemption models.Redemption) models.Redemption
	DeleteById(id int) error
	FindAll() []models.Redemption
	FindByUserID(id int) []models.Redemption
	FindByReward(id int) []models.Redemption
}
type redemptionRepository struct {
	BD *gorm.DB
}

// DeleteById implements RedemptionRepository.
func (r *redemptionRepository) DeleteById(id int) error {
	if result := r.BD.Delete(&models.Redemption{}, id).Error; result != nil {
		return result
	}
	return nil
}

// FindAll implements RedemptionRepository.
func (r *redemptionRepository) FindAll() []models.Redemption {
	var redemption []models.Redemption
	if result := r.BD.Find(&models.Redemption{}).Error; result != nil {
		return redemption
	}
	return redemption
}

// FindByReward implements RedemptionRepository.
func (r *redemptionRepository) FindByReward(id int) []models.Redemption {
	var redemption []models.Redemption
	if result := r.BD.Where("reward_id", id).Find(&models.Redemption{}).Error; result != nil {
		return redemption
	}
	return redemption
}

// FindByUserID implements RedemptionRepository.
func (r *redemptionRepository) FindByUserID(id int) []models.Redemption {
	var redemption []models.Redemption
	if result := r.BD.Where("user_id", id).Find(&models.Redemption{}).Error; result != nil {
		return redemption
	}
	return redemption
}

// Save implements RedemptionRepository.
func (r *redemptionRepository) Save(redemption models.Redemption) models.Redemption {
	r.BD.Save(&redemption)
	return redemption
}

// constructeur
func NewRedemptionRepository(repo *gorm.DB) RedemptionRepository {
	return &redemptionRepository{
		BD: repo,
	}
}
