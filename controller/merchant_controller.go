package controller

import (
	"loyalty-api/internal/models"
	service "loyalty-api/services"

	"github.com/gin-gonic/gin"
)

type MerchantController interface {
	SaveMerchant(ctx *gin.Context) models.Merchant
	FindAllMerchant() []models.Merchant
	DeleteMerchantById(id int) error
	FindMerchantById(id int) (models.Merchant, error)
	FindMerchantByEmail(email string) (models.Merchant, error)
	UpdateMerchant(id int, merchant models.Merchant) (models.Merchant, error)
}
type merchantController struct {
	controller service.MerchantService
}

// DeleteMerchantById implements MerchantController.
func (m *merchantController) DeleteMerchantById(id int) error {
	if err := m.controller.DeleteMerchantById(id); err != nil {
		return err
	}
	return nil
}

// FindAllMerchant implements MerchantController.
func (m *merchantController) FindAllMerchant() []models.Merchant {
	return m.controller.FindAllMerchant()
}

// FindMerchantByEmail implements MerchantController.
func (m *merchantController) FindMerchantByEmail(email string) (models.Merchant, error) {
	return m.controller.FindMerchantByEmail(email)
}

// FindMerchantById implements MerchantController.
func (m *merchantController) FindMerchantById(id int) (models.Merchant, error) {
	var merchant models.Merchant
	if merchant, err := m.controller.FindMerchantById(id); err != nil {
		return merchant, err
	}
	return merchant, nil
}

// SaveMerchant implements MerchantController.
func (m *merchantController) SaveMerchant(ctx *gin.Context) models.Merchant {
	var merchant models.Merchant
	ctx.BindJSON(&merchant)
	m.controller.SaveMerchant(merchant)
	return merchant
}

// UpdateMerchant implements MerchantController.
func (m *merchantController) UpdateMerchant(id int, merchant models.Merchant) (models.Merchant, error) {
	panic("unimplemented")
}

// contructeur
func NewMerchantController(service service.MerchantService) MerchantController {
	return &merchantController{
		controller: service,
	}
}
