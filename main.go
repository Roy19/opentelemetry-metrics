package main

import (
	"context"
	"log"
	"signoz-test/controllers"
	"signoz-test/db"
	"signoz-test/metrics"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	router := gin.Default()
	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping db", err)
	}
	defer db.Close()
	cleanup := metrics.InitMeterProvider(context.Background())
	defer cleanup(context.Background())
	router.Use(otelgin.Middleware("opentelemetry-demo"))
	router.POST("/cart", controllers.AddToCart)
	router.Run(":8080")
}
