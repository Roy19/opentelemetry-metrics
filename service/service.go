package service

import (
	"context"
	"errors"
	"log"
	"signoz-test/db"
	"signoz-test/db/generated"
	"signoz-test/dto"
)

func AddToCartService(ctx context.Context, cartItem dto.AddToCart) error {
	dbInstance := db.GetDBInstance()
	tx, err := dbInstance.BeginTx(ctx, nil)
	if err != nil {
		log.Println("[service.AddToCart] failed to start transaction", err)
		return errors.New("failed to start transaction")
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Println("[service.AddToCart] failed to successfully rollback transaction", err)
		}
	}()
	queries := generated.New(dbInstance).WithTx(tx)
	if err := queries.AddItemToCart(ctx, generated.AddItemToCartParams{
		CartID: int32(cartItem.CartId),
		Name:   cartItem.ItemName,
	}); err != nil {
		log.Println("[service.AddToCart] failed to add item to cart", err)
		return errors.New("failed to add to cart in db")
	}
	if err := tx.Commit(); err != nil {
		log.Println("[service.AddToCart] failed to commit transaction", err)
		return errors.New("failed to commit transaction")
	}
	return nil
}
