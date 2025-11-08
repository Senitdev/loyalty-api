package controller

import (
	"loyalty-api/internal/models"
	service "loyalty-api/services"

	"github.com/gin-gonic/gin"
)

type RedemptionController interface {
	Save(ctx *gin.Context) models.Redemption
	DeleteById(id int) error
	FindAll() []models.Redemption
	FindByUserID(id int) []models.Redemption
	FindByReward(id int) []models.Redemption
}
type redemptionController struct {
	controller service.RedemptionService
}

// DeleteById implements RedemptionController.
func (r *redemptionController) DeleteById(id int) error {
	if err := r.controller.DeleteById(id); err != nil {
		return err
	}
	return nil
}

// FindAll implements RedemptionController.
func (r *redemptionController) FindAll() []models.Redemption {
	return r.controller.FindAll()
}

// FindByReward implements RedemptionController.
func (r *redemptionController) FindByReward(id int) []models.Redemption {
	return r.controller.FindByReward(id)
}

// FindByUserID implements RedemptionController.
func (r *redemptionController) FindByUserID(id int) []models.Redemption {
	return r.controller.FindByUserID(id)
}

// Save implements RedemptionController.
func (r *redemptionController) Save(ctx *gin.Context) models.Redemption {
	var redemption models.Redemption
	ctx.BindJSON(&redemption)
	r.controller.Save(redemption)
	return redemption
}

// constructeur
func NewRedemptionController(service service.RedemptionService) RedemptionController {
	return &redemptionController{
		controller: service,
	}
}
