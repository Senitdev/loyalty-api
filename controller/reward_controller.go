package controller

import (
	"loyalty-api/internal/models"
	service "loyalty-api/services"

	"github.com/gin-gonic/gin"
)

type RewardController interface {
	Save(ctx *gin.Context) models.Reward
	DeleteById(id int) error
	FindAll() []models.Reward
	GetRewardByMerchant(id int) []models.Reward
	GetRewardById(id int) models.Reward
	UpdateReward(ctx *gin.Context, id int) models.Reward
}
type rewardController struct {
	controller service.RewardService
}

// GetRewardById implements RewardController.
func (r *rewardController) GetRewardById(id int) models.Reward {
	var reward = r.controller.GetRewardsById(id)
	return reward
}

// UpdateReward implements RewardController.
func (r *rewardController) UpdateReward(ctx *gin.Context, id int) models.Reward {
	var reward models.Reward
	ctx.BindJSON(&reward)
	r.controller.UpdateReward(id, reward)
	return reward
}

// DeleteById implements RewardController.
func (r *rewardController) DeleteById(id int) error {
	if err := r.controller.DeleteById(id); err != nil {
		return err
	}
	return nil
}

// FindAll implements RewardController.
func (r *rewardController) FindAll() []models.Reward {
	return r.controller.FindAll()
}

// GetRewardByMerchant implements RewardController.
func (r *rewardController) GetRewardByMerchant(id int) []models.Reward {
	return r.controller.GetRewardByMerchant(id)
}

// Save implements RewardController.
func (r *rewardController) Save(ctx *gin.Context) models.Reward {
	var reward models.Reward
	ctx.BindJSON(&reward)
	r.controller.Save(reward)
	return reward
}

// Contructeur
func NewRewardController(service service.RewardService) RewardController {
	return &rewardController{
		controller: service,
	}
}
