package http

import (
	_ "local/order-service/docs" // Swagger docs.
	"local/order-service/internal/app/config"
	appConfig "local/order-service/internal/app/config"
	"local/order-service/internal/controller/http/middleware"
	"local/order-service/pkg/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type Router struct {
	Engine         *gin.Engine
	userHandler    userHandler
	productHandler productHandler
	orderHandler   orderHandler
	config         config.Config
	apiGroup       *gin.RouterGroup
}

func (rt *Router) WithUserHandler(handler userHandler) *Router {
	rt.userHandler = handler

	users := rt.apiGroup.Group("/users")
	users.POST("", rt.userHandler.RegisterUser)
	users.GET("/:id", rt.userHandler.GetUser)
	users.PATCH("/:id", rt.userHandler.UpdateUser)
	users.DELETE("/:id", rt.userHandler.DeleteUser)

	return rt
}

func (rt *Router) WithProductHandler(handler productHandler) *Router {
	rt.productHandler = handler

	products := rt.apiGroup.Group("/products")
	products.POST("", rt.productHandler.AddProduct)
	products.GET("/:id", rt.productHandler.GetProduct)
	products.PATCH("/:id", rt.productHandler.UpdateProduct)
	products.DELETE("/:id", rt.productHandler.DeleteProduct)

	return rt
}

func (rt *Router) WithOrderHandler(handler orderHandler) *Router {
	rt.orderHandler = handler

	orders := rt.apiGroup.Group("/orders")
	orders.POST("", rt.orderHandler.CreateOrder)
	orders.GET("/:id", rt.orderHandler.GetOrder)

	return rt
}

func (rt *Router) WithLogger(l logger.Interface) *Router {
	rt.apiGroup.Use(middleware.Logger(l))
	return rt
}

// NewRouter -.
// Swagger spec:
// @title       Order Service  API
// @version     1.0
// @BasePath    /public/v1
func NewRouter(cfg *appConfig.Config, options ...func(*Router)) *Router {
	rt := &Router{}
	for _, opt := range options {
		opt(rt)
	}

	r := gin.Default()
	rt.Engine = r

	internal := r.Group("/internal")
	// Swagger
	if cfg.Swagger.Enabled {
		internal.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	if cfg.Tracing.Enabled {
		r.Use(otelgin.Middleware("order-service"))
	}

	api := r.Group("/public/api")
	v1 := api.Group("/v1")
	rt.apiGroup = v1

	return rt
}
