package controllers

import (
	"log"
	"signoz-test/dto"
	"signoz-test/metrics"
	"signoz-test/service"
	"time"

	"github.com/gin-gonic/gin"
)

func constructResponse(c *gin.Context, err error) {
	ctx := c.Request.Context()
	if err != nil {
		metrics.IncFailedRequests(ctx, 400)
		c.JSON(400, dto.Response{
			ErrorCode: 400,
			Message:   err.Error(),
		})
	} else {
		metrics.IncSuccessfulRequests(ctx, 201)
		c.JSON(201, dto.Response{
			Message: "operation success",
		})
	}
}

func AddToCart(c *gin.Context) {
	startTime := time.Now()
	var cartItem dto.AddToCart
	if err := c.BindJSON(&cartItem); err != nil {
		log.Println("[controller.AddToCart] failed to convert request to json", err)
		constructResponse(c, err)
		return
	}
	if err := cartItem.Validate(); err != nil {
		log.Println("[controller.AddToCart] failed to validate request", err)
		constructResponse(c, err)
		return
	}
	ctx := c.Request.Context()
	constructResponse(c, service.AddToCartService(ctx, cartItem))
	duration := time.Since(startTime)
	metrics.RecordLatency(ctx,
		float64(duration.Milliseconds()),
		c.Request.Method,
		"/cart",
	)
}
