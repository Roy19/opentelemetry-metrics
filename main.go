package main

import (
	"log"
	"signoz-test/controllers"
	"signoz-test/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping db", err)
	}
	router.POST("/cart", controllers.AddToCart)
	router.Run(":8080")
}
