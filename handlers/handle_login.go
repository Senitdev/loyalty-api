package handlers

import (
	"loyalty-api/controller"
	"loyalty-api/repository"
	service "loyalty-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Endpoint for login
func ParamLogin(ctx *gin.Engine, db *gorm.DB) {
	loginrepo := repository.NewLoginRepository(db)
	loginservice := service.NewLoginService(loginrepo)
	jwtService := service.NewJWTService()
	logincontroller := controller.NewLoginController(loginservice, jwtService)
	r := ctx.Group("/api/v1")
	r.POST("/auth", func(ctx *gin.Context) {
		token := logincontroller.Login(ctx)
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed"})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
		}
	})
}
