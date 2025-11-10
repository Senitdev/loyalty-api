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

func ParamTransactionRoutes(cx *gin.Engine, db *gorm.DB) {
	transactionRepo := repository.NewTransctionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionController := controller.NewTransactionController(transactionService)
	r := cx.Group("/api/v1")
	r.POST("/transaction", func(ctx *gin.Context) {
		ctx.JSON(200, transactionController.Save(ctx))
	})
	//Find ALL
	r.GET("/transaction", func(ctx *gin.Context) {
		ctx.JSON(200, transactionController.FindAll())
	})
	//Delete by ID
	r.DELETE("/transaction/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"bad request": ids})
		}
		transactionController.DeleteById(id)
		ctx.JSON(http.StatusNoContent, gin.H{"Delete Succes": id})
	})
	//Get Transaction by Merchant
	r.GET("/transaction/merchant/:merchantId/:startDate/:endDate", func(ctx *gin.Context) {
		startDate := ctx.Param("startDate")
		endDate := ctx.Param("endDate")
		idsMerchantId := ctx.Param("merchantId")
		merchantid, err := strconv.Atoi(idsMerchantId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Erreur": "Format id incorrect ou manque le id"})
			return
		}
		transactions, err := transactionController.FindbyMerchant(merchantid, startDate, endDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Erreur": "Bad request"})
			return
		}
		ctx.JSON(200, transactions)
	})
}
