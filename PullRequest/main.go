package main

import (
	"github_wb/infrastructure"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	router := gin.Default()

	port := os.Getenv("PORT")

	infrastructure.Routes(router)

	if port == "" {
		port = "8081"
	}

	router.Run(":" + port)

}
