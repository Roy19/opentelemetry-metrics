package interfaces

import (
	"context"
	"signoz-test/dto"
)

//go:generate mockgen -destination=../mocks/mock_service.go -package=mocks signoz-test/interfaces Service
type Service interface {
	AddItemToCart(ctx context.Context, cartItem dto.AddToCart) error
	GetItemsInCart(ctx context.Context, cartName string) (*dto.ItemsInCart, error)
}
