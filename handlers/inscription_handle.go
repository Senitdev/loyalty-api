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
	clientsRepo := repository.NewClientsRepository(db)
	clientsService := service.NewClientsService(clientsRepo)
	clientsController := controller.NewClientsController(clientsService)
	r := cx.Group("/api/v1")
	//Save
	r.POST("/inscription", func(ctx *gin.Context) {
		ctx.JSON(200, clientsController.Save(ctx))
	})

}
