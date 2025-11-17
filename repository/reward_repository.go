package repository

import (
	"loyalty-api/internal/models"

	"gorm.io/gorm"
)

type RewardRepository interface {
	Save(reward models.Reward) models.Reward
	DeleteById(id int) error
	FindAll() []models.Reward
	GetRewardByMerchant(id int) []models.Reward
	GetRewardById(id int) models.Reward
	UpdateReward(id int, reward models.Reward) models.Reward
}
type rewardRepository struct {
	DB *gorm.DB
}

// GetRewardById implements RewardRepository.
func (r *rewardRepository) GetRewardById(id int) models.Reward {
	var reward models.Reward
	if result := r.DB.Find(&reward, id).Error; result != nil {
		return reward
	}
	return reward
}

// UpdateReward implements RewardRepository.
func (r *rewardRepository) UpdateReward(id int, reward models.Reward) models.Reward {
	var rewardExist models.Reward
	if result := r.DB.Find(&rewardExist, id).Error; result != nil {
		return rewardExist
	}
	reward.CreatedAt = rewardExist.CreatedAt
	reward.MerchantID = rewardExist.MerchantID
	reward.Merchant = rewardExist.Merchant
	reward.ID = rewardExist.ID
	if reward.Description == "" {
		reward.Description = rewardExist.Description
	}
	if reward.Title == "" {
		reward.Title = rewardExist.Title
	}
	if reward.PointsRequired < 1 {
		reward.PointsRequired = rewardExist.PointsRequired
	}
	if resultat := r.DB.Save(reward).Error; resultat != nil {
		return reward
	}
	return reward
}

// DeleteById implements RewardRepository.
func (r *rewardRepository) DeleteById(id int) error {
	result := r.DB.Delete(&models.Reward{}, id)
	if result != nil {
		return result.Error
	}
	return nil
}

// FindAll implements RewardRepository.
func (r *rewardRepository) FindAll() []models.Reward {
	var reward []models.Reward
	result := r.DB.Find(&reward)
	if result != nil {
		return reward
	}
	return reward
}

// GetRewardByMerchant implements RewardRepository.
func (r *rewardRepository) GetRewardByMerchant(id int) []models.Reward {
	var reward []models.Reward
	if result := r.DB.Where("merchant_id = ?", id).Find(&reward).Error; result != nil {
		return reward
	}
	return reward
}

// Save implements RewardRepository.
func (r *rewardRepository) Save(reward models.Reward) models.Reward {
	r.DB.Save(&reward)
	return reward
}

// CONSTRUCTEUR
func NewRewardRepository(repo *gorm.DB) RewardRepository {
	return &rewardRepository{
		DB: repo,
	}
}
