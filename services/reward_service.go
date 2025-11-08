package service

import (
	"loyalty-api/internal/models"
	"loyalty-api/repository"
	"time"
)

type RewardService interface {
	Save(reward models.Reward) models.Reward
	DeleteById(id int) error
	FindAll() []models.Reward
	GetRewardByMerchant(id int) []models.Reward
	GetRewardsById(id int) models.Reward
	UpdateReward(id int, reward models.Reward) models.Reward
}
type rewardService struct {
	service repository.RewardRepository
}

// GetRewardById implements RewardService.
func (r *rewardService) GetRewardsById(id int) models.Reward {
	var reward = r.service.GetRewardById(id)
	return reward
}

// UpdateReward implements RewardService.
func (r *rewardService) UpdateReward(id int, reward models.Reward) models.Reward {
	return r.service.UpdateReward(id, reward)
}

// DeleteById implements RewardService.
func (r *rewardService) DeleteById(id int) error {
	if err := r.service.DeleteById(id); err != nil {
		return err
	}
	return nil
}

// FindAll implements RewardService.
func (r *rewardService) FindAll() []models.Reward {
	return r.service.FindAll()
}

// GetRewardByMerchant implements RewardService.
func (r *rewardService) GetRewardByMerchant(id int) []models.Reward {
	return r.service.GetRewardByMerchant(id)
}

// Save implements RewardService.
func (r *rewardService) Save(reward models.Reward) models.Reward {
	reward.CreatedAt = time.Now()
	return r.service.Save(reward)
}

// constructeur
func NewRewardService(repo repository.RewardRepository) RewardService {
	return &rewardService{
		service: repo,
	}
}
