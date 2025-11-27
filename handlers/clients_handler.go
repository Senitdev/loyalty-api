package handlers

import (
	"loyalty-api/controller"
	"loyalty-api/internal/models"
	"loyalty-api/repository"
	service "loyalty-api/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParamClientsRoutes(cx *gin.Engine, db *gorm.DB) {
	//declaration des variables
	clientsRepo := repository.NewClientsRepository(db)
	clientsService := service.NewClientsService(clientsRepo)
	clientsController := controller.NewClientsController(clientsService)
	r := cx.Group("/api/v1")

	//GET ALL
	r.GET("/clients", func(ctx *gin.Context) {
		ctx.JSON(200, clientsController.FindAll())
	})
	//Get By ID
	r.GET("/clients/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		if ids == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"id incorrect ou manaquant": ids})
		}
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"id incorrect ou manaquant": ids})
		}
		ctx.JSON(200, clientsController.GetClientById(id))
	})
	//delete clients by ID
	r.DELETE("/clients/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		if ids == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"id incorrect ou manaquant": ids})
		}
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"id incorrect ou manaquant": ids})
		}
		ctx.JSON(http.StatusNoContent, clientsController.DeleteClientsById(id))
	})
	//update
	r.PUT("/clients/:id", func(ctx *gin.Context) {
		ids := ctx.Param("id")
		if ids == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"id incorrect ou manaquant": ids})
		}
		id, err := strconv.Atoi(ids)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"id incorrect ou manaquant": ids})
		}
		ctx.JSON(http.StatusNoContent, clientsController.UpdateClient(ctx, id))
	})
	//Search client BY Mobile OR Email
	r.GET("/clients/search", func(ctx *gin.Context) {
		query := strings.TrimSpace(ctx.Query("query"))

		if query == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Paramètre 'query' requis (email ou téléphone)"})
			return
		}
		var client models.Clients
		client, err := clientsController.SearchClient(query)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Client non trouve"})
			return
		}
		ctx.JSON(200, client)
	})

}
