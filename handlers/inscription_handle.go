package handlers

import (
	"loyalty-api/controller"
	"loyalty-api/repository"
	service "loyalty-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParamInscriptionRoutes(cx *gin.Engine, db *gorm.DB) {
	//declaration des variables
	//Inscription routes
	//clients
	clientsRepo := repository.NewClientsRepository(db)
	clientsService := service.NewClientsService(clientsRepo)
	clientsController := controller.NewClientsController(clientsService)
	//Merchant inscription
	merchantRepo := repository.NewMerchantRepository(db)
	merchandService := service.NewMerchantService(merchantRepo)
	merchantController := controller.NewMerchantController(merchandService)
	r := cx.Group("/api/v1")
	//Save
	r.POST("/inscription/client", func(ctx *gin.Context) {
		ctx.JSON(200, clientsController.Save(ctx))
	})
	//Merchant
	r.POST("/inscription/merchant", func(ctx *gin.Context) {
		ctx.JSON(200, merchantController.SaveMerchant(ctx))
	})
}
