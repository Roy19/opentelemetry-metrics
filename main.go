package main

import (
	"context"
	"log"
	"os"
	"signoz-test/controllers"
	"signoz-test/db"
	"signoz-test/metrics"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

var serviceName = os.Getenv("SERVICE_NAME")

func main() {
	router := gin.Default()
	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping db", err)
	}
	defer db.Close()
	cleanup := metrics.InitMeterProvider(context.Background())
	defer cleanup(context.Background())
	router.Use(otelgin.Middleware(serviceName))
	cartGroup := router.Group("/cart")
	cartGroup.POST("/add", controllers.AddToCart)
	cartGroup.GET("/:cartName", controllers.GetItemsInCart)
	router.Run(":8080")
}
