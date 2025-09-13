package main

import (
	"context"
	"log"
	"os"
	"signoz-test/controllers"
	"signoz-test/db"
	"signoz-test/interfaces"
	"signoz-test/metrics"
	"signoz-test/service"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

var serviceName = os.Getenv("SERVICE_NAME")

func initDependencies() interfaces.Controller {
	controller := controllers.NewController(
		service.NewService(
			db.GetDBInstance(),
		),
	)
	return controller
}

func main() {
	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping db", err)
	}
	defer db.Close()

	cleanupMetricsProvider := metrics.InitMeterProvider(context.Background())
	defer cleanupMetricsProvider(context.Background())

	controller := initDependencies()

	router := gin.Default()
	router.Use(otelgin.Middleware(serviceName))
	cartGroup := router.Group("/cart")
	cartGroup.POST("/add", controller.AddToCart)
	cartGroup.GET("/:cartName", controller.GetItemsInCart)

	router.Run(":8080")
}
