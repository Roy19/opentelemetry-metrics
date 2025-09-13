package controllers

import (
	"log"
	"signoz-test/dto"
	"signoz-test/metrics"
	"signoz-test/service"

	"github.com/gin-gonic/gin"
)

func constructErrorResponse(c *gin.Context, err error) {
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
	var cartItem dto.AddToCart
	if err := c.BindJSON(&cartItem); err != nil {
		log.Println("failed to convert request to json", err)
		constructErrorResponse(c, err)
		return
	}
	ctx := c.Request.Context()
	err := service.AddToCartService(ctx, cartItem)
	constructErrorResponse(c, err)
}
