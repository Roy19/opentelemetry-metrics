package controllers

import (
	"log"
	"signoz-test/dto"
	"signoz-test/metrics"
	"signoz-test/service"
	"time"

	"github.com/gin-gonic/gin"
)

func constructResponse(c *gin.Context, err error, message interface{}) {
	ctx := c.Request.Context()
	if err != nil {
		metrics.IncFailedRequests(ctx, 400)
		c.JSON(400, dto.Response{
			Code:    400,
			Message: err.Error(),
		})
	} else {
		if message != nil {
			metrics.IncSuccessfulRequests(ctx, 200)
			c.JSON(200, dto.Response{
				Code:    200,
				Message: message,
			})
		} else {
			metrics.IncSuccessfulRequests(ctx, 201)
			c.JSON(201, dto.Response{
				Code:    201,
				Message: "operation_successful",
			})
		}
	}
}

func AddToCart(c *gin.Context) {
	startTime := time.Now()
	var cartItem dto.AddToCart
	if err := c.BindJSON(&cartItem); err != nil {
		log.Println("[controller.AddToCart] failed to convert request to json", err)
		constructResponse(c, err, nil)
		return
	}
	if err := cartItem.Validate(); err != nil {
		log.Println("[controller.AddToCart] failed to validate request", err)
		constructResponse(c, err, nil)
		return
	}
	ctx := c.Request.Context()
	constructResponse(c, service.AddToCartService(ctx, cartItem), nil)
	duration := time.Since(startTime)
	metrics.RecordLatency(ctx,
		float64(duration.Milliseconds()),
		c.Request.Method,
		"/cart/add",
	)
}

func GetItemsInCart(c *gin.Context) {
	startTime := time.Now()
	ctx := c.Request.Context()
	cartName := c.Param("cartName")
	items, err := service.GetItemsFromCart(ctx, cartName)
	if err != nil {
		log.Println("[controller.GetItemsInCart] failed to get items in cart", err)
		constructResponse(c, err, nil)
		return
	}
	constructResponse(c, nil, items)
	duration := time.Since(startTime)
	metrics.RecordLatency(ctx,
		float64(duration.Milliseconds()),
		c.Request.Method,
		"/cart/:cartName",
	)
}
