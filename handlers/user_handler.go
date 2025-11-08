package handlers

import (
	"loyalty-api/controller"
	"loyalty-api/repository"
	service "loyalty-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParamUserRoutes(cx *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	r := cx.Group("/api/v1")
	r.POST("/user", func(ctx *gin.Context) {
		ctx.JSON(200, userController.SaveUser(ctx))
	})
}
