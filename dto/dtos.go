package dto

type AddToCart struct {
	CartId   int    `json:"cart_id"`
	ItemName string `json:"item_name"`
}

type Response struct {
	ErrorCode int    `json:"error_code,omitempty"`
	Message   string `json:"message"`
}
