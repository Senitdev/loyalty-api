package handlers

import (
	"loyalty-api/controller"
	midllewares "loyalty-api/middlewares"
	"loyalty-api/repository"
	service "loyalty-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParamRewardRoutes(cx *gin.Engine, db *gorm.DB) {
	rewardRepo := repository.NewRewardRepository(db)
	rewardService := service.NewRewardService(rewardRepo)
	rewardController := controller.NewRewardController(rewardService)
	r := cx.Group("/api/v1", midllewares.AuthorizeJWT())
	r.POST("/reward", func(ctx *gin.Context) {
		ctx.JSON(200, rewardController.Save(ctx))
	})
	//Get ALL Reward
	r.GET("/reward", func(ctx *gin.Context) {
		ctx.JSON(200, rewardController.FindAll())
	})
	//Delete reward
	r.DELETE("/reward/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Id inconnu ou indefinii": id})
		}
		rewardController.DeleteById(id)
		ctx.JSON(http.StatusNoContent, gin.H{"Succes delete": id})
	})
	//Get reward by merchant
	r.GET("/reward/merchant/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Id inconnu ou no defini": id})
		}
		ctx.JSON(200, rewardController.GetRewardByMerchant(id))
	})
	//get reward by ID
	r.GET("/reward/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Id inconnu ou no defini": id})
		}
		ctx.JSON(200, rewardController.GetRewardById(id))
	})
	//Update Reward by ID
	r.PUT("/reward/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Id inconnu ou no defini": id})
		}
		ctx.JSON(200, rewardController.UpdateReward(ctx, id))
	})
}
