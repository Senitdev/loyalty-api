package handlers

import (
	"loyalty-api/controller"
	"loyalty-api/repository"
	service "loyalty-api/services"

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
	r.POST("/auth", logincontroller.Login)
}
