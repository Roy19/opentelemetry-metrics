package interfaces

import (
	"context"
	"signoz-test/dto"
)

type Service interface {
	AddItemToCart(ctx context.Context, cartItem dto.AddToCart) error
	GetItemsInCart(ctx context.Context, cartName string) (*dto.ItemsInCart, error)
}
