package handlers

import (
	"loyalty-api/controller"
	"loyalty-api/repository"
	service "loyalty-api/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParamRewardRoutes(cx *gin.Engine, db *gorm.DB) {
	rewardRepo := repository.NewRewardRepository(db)
	rewardService := service.NewRewardService(rewardRepo)
	rewardController := controller.NewRewardController(rewardService)
	r := cx.Group("/api/v1")
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
	r.GET("/reward/merchant", func(ctx *gin.Context) {
		query := strings.TrimSpace(ctx.Query("query"))
		if query == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"Manque le email ": query})
			return
		}
		id, err := strconv.Atoi(query)
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
