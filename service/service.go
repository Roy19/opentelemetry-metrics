package service

import (
	"context"
	"database/sql"
	"log"
	"signoz-test/db/generated"
	"signoz-test/dto"
	"signoz-test/errors"
	"signoz-test/interfaces"
	"signoz-test/metrics"
)

type Service struct {
	dbInstance *sql.DB
}

func NewService(db *sql.DB) interfaces.Service {
	return &Service{
		dbInstance: db,
	}
}

func (s *Service) AddItemToCart(ctx context.Context, cartItem dto.AddToCart) error {
	tx, err := s.dbInstance.BeginTx(ctx, nil)
	if err != nil {
		log.Println("[service.AddToCart] failed to start transaction", err)
		return errors.ErrTxStart
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Println("[service.AddToCart] failed to successfully rollback transaction", err)
		}
	}()
	queries := generated.New(s.dbInstance).WithTx(tx)
	cartId, err := queries.GetCartIdGivenName(ctx, *cartItem.CartName)
	if err != nil {
		log.Println("failed to get cart name", err)
		return errors.ErrCartDoesNotExists
	}
	if err := queries.AddItemToCart(ctx, generated.AddItemToCartParams{
		CartID: cartId,
		Name:   *cartItem.ItemName,
	}); err != nil {
		log.Println("[service.AddToCart] failed to add item to cart", err)
		return errors.ErrFailedCartAdd
	}
	if err := tx.Commit(); err != nil {
		log.Println("[service.AddToCart] failed to commit transaction", err)
		return errors.ErrTxCommit
	}
	return nil
}

func (s *Service) GetItemsInCart(ctx context.Context, cartName string) (*dto.ItemsInCart, error) {
	tx, err := s.dbInstance.BeginTx(ctx, nil)
	if err != nil {
		log.Println("[service.GetItemsFromCart] failed to start transaction", err)
		return nil, errors.ErrTxStart
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Println("[service.GetItemsFromCart] failed to successfully rollback transaction", err)
		}
	}()
	queries := generated.New(s.dbInstance).WithTx(tx)
	id, err := queries.GetCartIdGivenName(ctx, cartName)
	if err != nil {
		log.Println("[service.GetItemsFromCart] failed to find cart with given name", err)
		return nil, errors.ErrCartDoesNotExists
	}
	items, err := queries.GetItemsInCart(ctx, id)
	if err != nil {
		log.Println("[service.GetItemsFromCart] failed to get items in cart", err)
		return nil, errors.ErrItemsGet
	}
	allItems := dto.ItemsInCart{
		Items: make([]string, 0),
	}
	for _, item := range items {
		allItems.Items = append(allItems.Items, item.Name)
	}
	metrics.RecordItemsInCart(ctx, cartName, len(allItems.Items))
	if err := tx.Commit(); err != nil {
		log.Println("[service.GetItemsFromCart] failed to commit transaction", err)
		return nil, err
	}
	return &allItems, nil
}
