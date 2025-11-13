package controller

import (
	"loyalty-api/controller/dto"
	"loyalty-api/internal/models"
	service "loyalty-api/services"

	"github.com/gin-gonic/gin"
)

type LoyaltyCardController interface {
	Save(ctx *gin.Context) models.LoyaltyCard
	DeleteById(id int) error
	FindAll() []models.LoyaltyCard
	FindAllByUser(id int) []models.Clients
	FindAllByMerchant(id int) []models.LoyaltyCard
	AddPointsCard(ctx *gin.Context, points int) (models.LoyaltyCard, error)
	SoldePointsClient(ctx *gin.Context) (int, error)
	RetraitPointsClient(ctx *gin.Context, points int) (models.LoyaltyCard, error)
	SoldeMerchant(merchantId int) int
	SoldePointsClientByAllMerchants(clientId int) (int, error)
	FindAllMerchantByClient(clientId int) []dto.MerchantDto
}
type loyaltyCardController struct {
	controller service.LoyaltyCardService
}

// FindAllMerchantByClient implements LoyaltyCardController.
// Trouver la liste des boutiques abonnee pour un client
func (l *loyaltyCardController) FindAllMerchantByClient(clientId int) []dto.MerchantDto {
	return l.controller.FindAllMerchantByClient(clientId)
}

// SoldePointsClientByAllMerchants implements LoyaltyCardController.
func (l *loyaltyCardController) SoldePointsClientByAllMerchants(clientId int) (int, error) {
	return l.controller.SoldePointsClientByAllMerchants(clientId)
}

// SoldeMerchant implements LoyaltyCardController.
func (l *loyaltyCardController) SoldeMerchant(merchantId int) int {
	return l.controller.SoldeMerchant(merchantId)
}

// RetraitPointsClient implements LoyaltyCardController.
func (l *loyaltyCardController) RetraitPointsClient(ctx *gin.Context, points int) (models.LoyaltyCard, error) {
	var loyalty models.LoyaltyCard
	ctx.BindJSON(&loyalty)
	loyalty, err := l.controller.RetraitPointsClient(loyalty, points)
	if err != nil {
		return loyalty, err
	}
	return loyalty, nil
}

// SoldePointsClient implements LoyaltyCardController.
func (l *loyaltyCardController) SoldePointsClient(ctx *gin.Context) (int, error) {
	var loyalcard models.LoyaltyCard
	ctx.BindJSON(&loyalcard)
	points, err := l.controller.SoldePointsClient(loyalcard)
	if err != nil {
		return points, err
	}
	return points, nil
}

// AddPointsCard implements LoyaltyCardController.
func (l *loyaltyCardController) AddPointsCard(ctx *gin.Context, points int) (models.LoyaltyCard, error) {
	var loyalty models.LoyaltyCard
	ctx.BindJSON(&loyalty)
	loyalty, err := l.controller.AddPointsCard(loyalty, points)
	if err != nil {
		return loyalty, err
	}
	return loyalty, nil
}

// DeleteById implements LoyaltyCardController.
func (l *loyaltyCardController) DeleteById(id int) error {
	if err := l.controller.DeleteById(id); err != nil {
		return err
	}
	return nil
}

// FindAll implements LoyaltyCardController.
func (l *loyaltyCardController) FindAll() []models.LoyaltyCard {
	return l.controller.FindAll()
}

// FindAllByMerchant implements LoyaltyCardController.
func (l *loyaltyCardController) FindAllByMerchant(id int) []models.LoyaltyCard {
	return l.controller.FindAllByMerchant(id)
}

// FindAllByUser implements LoyaltyCardController.
func (l *loyaltyCardController) FindAllByUser(id int) []models.Clients {
	return l.controller.FindAllByUser(id)
}

// Save implements LoyaltyCardController.
func (l *loyaltyCardController) Save(ctx *gin.Context) models.LoyaltyCard {
	var loyaltyCard models.LoyaltyCard
	ctx.BindJSON(&loyaltyCard)
	l.controller.Save(loyaltyCard)
	return loyaltyCard
}

// Contructeur
func NewLoyaltyCardController(service service.LoyaltyCardService) LoyaltyCardController {
	return &loyaltyCardController{
		controller: service,
	}
}
