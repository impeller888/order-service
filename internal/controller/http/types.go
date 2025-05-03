package http

import "github.com/gin-gonic/gin"

type userHandler interface {
	GetUser(c *gin.Context)
	RegisterUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type productHandler interface {
	GetProduct(c *gin.Context)
	AddProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
}

type orderHandler interface {
	GetOrder(c *gin.Context)
	CreateOrder(c *gin.Context)
	// GetUserOrders(c *gin.Context)
}
