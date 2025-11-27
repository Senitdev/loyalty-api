package handlers

import (
	"loyalty-api/controller"
	"loyalty-api/repository"
	service "loyalty-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParamRedemptionRoutes(cx *gin.Engine, db *gorm.DB) {
	redemptionRepo := repository.NewRedemptionRepository(db)
	redemptionService := service.NewRedemptionService(redemptionRepo)
	redemptionController := controller.NewRedemptionController(redemptionService)
	r := cx.Group("/api/v1")
	r.POST("/redemption", func(ctx *gin.Context) {
		ctx.JSON(200, redemptionController.Save(ctx))
	})
	r.DELETE("/redemption/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Id inconnu ou Manquant": id})
		}
		if err := redemptionController.DeleteById(id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"id incorrect ou inconnu": id})
			return
		}
		redemptionController.DeleteById(id)
		ctx.JSON(http.StatusNoContent, gin.H{"Succes delete": id})
	})
	//Get all
	r.GET("/redemption", func(ctx *gin.Context) {
		ctx.JSON(200, redemptionController.FindAll())
	})
	//Get All by user
	r.GET("/redemption/user/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ids, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"id inconnu ou incorrect": ids})
		}
		ctx.JSON(200, redemptionController.FindByUserID(ids))
	})
	//Get All by Reward
	r.GET("/redemption/reward/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"id inconnu ou manquant": id})
		}
		ctx.JSON(200, redemptionController.FindByReward(id))
	})

}
