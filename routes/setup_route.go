package routes

import (
	"loyalty-api/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	//config les Cors
	//cors config
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // ton front React/Next.js
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(config))
	//On defini les routes
	//login
	handlers.ParamLogin(r, db)
	//Users
	handlers.ParamUserRoutes(r, db)
	//Merchant
	handlers.ParamMerchantRoutes(r, db)
	//LoyaltyCard
	handlers.ParamLoyaltyCardRoutes(r, db)
	//Redemption
	handlers.ParamRedemptionRoutes(r, db)
	//Reward
	handlers.ParamRewardRoutes(r, db)
	//transction
	handlers.ParamTransactionRoutes(r, db)
	//Clients
	handlers.ParamClientsRoutes(r, db)
	return r
}
