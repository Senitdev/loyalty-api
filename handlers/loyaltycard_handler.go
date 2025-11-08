package handlers

import (
	"loyalty-api/controller"
	"loyalty-api/controller/dto"
	"loyalty-api/internal/models"
	"loyalty-api/repository"
	service "loyalty-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParamLoyaltyCardRoutes(cx *gin.Engine, db *gorm.DB) {
	loyaltycardRepo := repository.NewLoyaltyCardRepository(db)
	loyaltycarService := service.NewLoyaltyCardService(loyaltycardRepo)
	loyaltyCardController := controller.NewLoyaltyCardController(loyaltycarService)
	r := cx.Group("/api/v1")
	r.POST("/loyaltycard", func(ctx *gin.Context) {
		ctx.JSON(200, loyaltyCardController.Save(ctx))
	})
	r.GET("/loyaltycard", func(ctx *gin.Context) {
		ctx.JSON(200, loyaltyCardController.FindAll())
	})
	r.GET("/loyaltycard/user/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Manque ID ou inconnu": id})
		}

		ctx.JSON(200, loyaltyCardController.FindAllByUser(id))
	})
	//Get By user
	r.GET("/loyaltycard/merchand/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"ID inconnu ou manquant": id})
		}
		ctx.JSON(200, loyaltyCardController.FindAllByMerchant(id))
	})
	//Delete By ID
	r.DELETE("/loyaltycard/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Id inconnu ou invalide": ids})
		}
		ctx.JSON(200, loyaltyCardController.DeleteById(id))
	})
	//Add Points
	r.POST("/loyaltycard/:points", func(ctx *gin.Context) {
		pointStr := ctx.Param("points")
		points, err := strconv.Atoi(pointStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Erreur": "manque le nbre de points"})
			return
		}
		if points < 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{"Erreur": "Nbres de points doit etre positive"})
			return
		}
		var loyalcard models.LoyaltyCard
		loyalcard, err = loyaltyCardController.AddPointsCard(ctx, points)
		ctx.JSON(200, loyalcard)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Erreur": "Mauvaise format"})
			//return
		}
	})
	//Solde loyalty
	r.POST("/loyaltycard/solde", func(ctx *gin.Context) {
		points, err := loyaltyCardController.SoldePointsClient(ctx)
		var loyalcard models.LoyaltyCard
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Aucune enregistrement": points})
			return
		}
		loyalcard.Points = points
		ctx.JSON(200, loyalcard)

	})
	//Retrait loyaltycard
	r.POST("/loyaltycard/retrait/:points", func(ctx *gin.Context) {
		pointStr := ctx.Param("points")
		points, err := strconv.Atoi(pointStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Erreur": "manque le nbre de points"})
			return
		}
		if points < 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{"Erreur": "Nbres de points doit etre positive"})
			return
		}
		var loyalcard models.LoyaltyCard
		loyalcard, err = loyaltyCardController.RetraitPointsClient(ctx, points)
		ctx.JSON(200, loyalcard)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Erreur": "Mauvaise format"})
			//return
		}
	})
	//Solde
	r.GET("/loyaltycard/merchant/solde/:merchantId", func(ctx *gin.Context) {
		ids := ctx.Param("merchantId")
		Idmerchant, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"Erreur": "Manque le Id merchant"})
			return
		}
		var solde = loyaltyCardController.SoldeMerchant(Idmerchant)
		var loyaltyDto dto.LoyaltycardDto
		loyaltyDto.Merchant = Idmerchant
		loyaltyDto.Solde = solde
		ctx.JSON(200, loyaltyDto)
	})
}
