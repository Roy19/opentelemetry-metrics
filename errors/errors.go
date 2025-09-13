package errors

import "errors"

var (
	ErrTxStart           = errors.New("failed to start transaction")
	ErrCartDoesNotExists = errors.New("cart with given name does not exists")
	ErrTxCommit          = errors.New("failed to commit transaction")
	ErrFailedCartAdd     = errors.New("failed to add item to cart")
	ErrItemsGet          = errors.New("failed to get items in cart")
)
