package handlers

import (
	"loyalty-api/controller"
	"loyalty-api/internal/dto"
	"loyalty-api/repository"
	service "loyalty-api/services"
	"net/http"
	"strings"

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
	//Get user by Email
	r.GET("/user/search", func(ctx *gin.Context) {
		query := strings.TrimSpace(ctx.Query("query"))
		if query == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"Manque le email ": query})
			return
		}
		var userDTO dto.UserDTO
		userDTO, err := userController.GetUserByEmail(query)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Erreur": "err"})
			return
		}
		ctx.JSON(http.StatusOK, userDTO)
	})
}
