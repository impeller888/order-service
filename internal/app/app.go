// Package app configures and runs application.
package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"local/order-service/internal/app/config"

	v1 "local/order-service/internal/controller/http"

	orderHandler "local/order-service/internal/controller/http/v1/handlers/order"
	productHandler "local/order-service/internal/controller/http/v1/handlers/product"
	userHandler "local/order-service/internal/controller/http/v1/handlers/user"
	orderRepo "local/order-service/internal/repo/db/order"
	productRepo "local/order-service/internal/repo/db/product"
	userRepo "local/order-service/internal/repo/db/user"
	orderSrv "local/order-service/internal/services/order"
	productSrv "local/order-service/internal/services/product"
	userSrv "local/order-service/internal/services/user"

	"local/order-service/pkg/logger"
	"local/order-service/pkg/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MakeRouter(cfg *config.Config, pool *pgxpool.Pool, lg logger.Interface) *v1.Router {
	userRepository := userRepo.NewUserRepository(pool)
	productRepository := productRepo.NewProductRepository(pool)
	orderRepository := orderRepo.NewOrderRepository(pool)
	// Use case
	userService := userSrv.NewUserService(userRepository)
	productService := productSrv.NewProductService(productRepository)
	orderService := orderSrv.NewOrderService(userService, orderRepository, productRepository)

	userHandler := userHandler.NewHandler(userService)
	productHandler := productHandler.NewHandler(productService)
	orderHandler := orderHandler.NewHandler(orderService)

	return v1.NewRouter(cfg).
		WithLogger(lg).
		WithUserHandler(userHandler).
		WithProductHandler(productHandler).
		WithOrderHandler(orderHandler)
}

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	lg := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.Connect(cfg.DB.PG, lg)
	if err != nil {
		lg.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	router := MakeRouter(cfg, pg.Pool, lg)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router.Engine,
	}

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("Stopped serving new connections.")
	}()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete.")
}
