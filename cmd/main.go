package main

import (
	"fmt"
	"loyalty-api/database"
	"loyalty-api/routes"
	"os"
)

func main() {
	database.Connect()
	//server.SetUpRoute(db.DB)
	r := routes.SetupRoutes(database.DB)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8090"
	}
	r.Run(":" + port)
	fmt.Printf("Server Running at port %s", port)
}
