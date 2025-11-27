package handlers

import (
	"loyalty-api/controller"
	"loyalty-api/internal/models"
	"loyalty-api/repository"
	service "loyalty-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParamMerchantRoutes(cx *gin.Engine, db *gorm.DB) {
	merchantRepo := repository.NewMerchantRepository(db)
	merchandService := service.NewMerchantService(merchantRepo)
	merchantController := controller.NewMerchantController(merchandService)
	r := cx.Group("/api/v1")
	r.GET("/merchant", func(ctx *gin.Context) {
		ctx.JSON(200, merchantController.FindAllMerchant())
	})
	r.GET("merchant/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Erreur id": id})
		}
		var merchant models.Merchant
		merchant, err = merchantController.FindMerchantById(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Une erreur est survenu": ids})
		}
		ctx.JSON(200, merchant)
	})

	r.GET("/merchant/email/:email", func(ctx *gin.Context) {
		email := ctx.Param("email")
		if email == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"Manque email": email})
		}
		var merchant models.Merchant
		merchant, err := merchantController.FindMerchantByEmail(email)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Bad request": email})
		}
		ctx.JSON(200, merchant)
	})
	r.DELETE("/merchant/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Manque l'Id": id})
		}
		if err := merchantController.DeleteMerchantById(id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"aucun enregistrement": id})
		}
		ctx.JSON(http.StatusNoContent, gin.H{"Succes delete": id})

	})
}
