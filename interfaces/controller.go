package interfaces

import "github.com/gin-gonic/gin"

type Controller interface {
	AddToCart(c *gin.Context)
	GetItemsInCart(c *gin.Context)
}
