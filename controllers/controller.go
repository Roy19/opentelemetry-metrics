package controllers

import (
	"errors"
	"log"
	"signoz-test/db"
	"signoz-test/db/generated"
	"signoz-test/dto"
	"signoz-test/metrics"

	"github.com/gin-gonic/gin"
)

func constructErrorResponse(c *gin.Context, err error, code int) {
	ctx := c.Request.Context()
	if err != nil {
		metrics.IncFailedRequests(ctx, code)
		c.JSON(code, dto.Response{
			ErrorCode: code,
			Message:   err.Error(),
		})
	} else {
		metrics.IncSuccessfulRequests(ctx, code)
		c.JSON(code, dto.Response{
			Message: "operation success",
		})
	}
}

func AddToCart(c *gin.Context) {
	var addToCart dto.AddToCart
	if err := c.BindJSON(&addToCart); err != nil {
		log.Println("failed to convert request to json", err)
		constructErrorResponse(c, err, 400)
		return
	}
	ctx := c.Request.Context()
	dbInstance := db.GetDBInstance()
	tx, err := dbInstance.BeginTx(ctx, nil)
	if err != nil {
		log.Println("failed to start a transaction", err)
		constructErrorResponse(c, err, 500)
		return
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Println("failed to successfully rollback transaction", err)
		}
	}()
	queries := generated.New(dbInstance).WithTx(tx)
	if err := queries.AddItemToCart(ctx, generated.AddItemToCartParams{
		CartID: int32(addToCart.CartId),
		Name:   addToCart.ItemName,
	}); err != nil {
		log.Println("failed to add item to cart", err)
		constructErrorResponse(c, errors.New("failed to add item to cart"), 500)
		return
	}
	if err := tx.Commit(); err != nil {
		log.Println("failed to commit transaction", err)
		constructErrorResponse(c, errors.New("failed to add commit transaction"), 500)
		return
	}
	constructErrorResponse(c, nil, 200)
}
