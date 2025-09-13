package dto

import "errors"

type AddToCart struct {
	CartName *string `json:"cart_name,omitempty"`
	ItemName *string `json:"item_name,omitempty"`
}

func (a *AddToCart) Validate() error {
	if a.CartName == nil {
		return errors.New("cart_name cannot be empty")
	}
	if a.ItemName == nil {
		return errors.New("item_name cannot be empty")
	}
	return nil
}

type Response struct {
	ErrorCode int    `json:"error_code,omitempty"`
	Message   string `json:"message"`
}
