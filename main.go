package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"bank-api/src/config"
	"bank-api/src/controllers"
	"bank-api/src/middlewares"
	"bank-api/src/repositories"
	"bank-api/src/services"
)

func main() {
	// load the configuration
	cfg := config.LoadConfig()

	userRepo := repositories.NewUserRepository(cfg.DataPath)
	paymentRepo := repositories.NewPaymentRepository(cfg.DataPath)

	userService := services.NewUserService(userRepo)
	paymentService := services.NewPaymentService(paymentRepo, userRepo)

	userController := controllers.NewUserController(userService)
	paymentController := controllers.NewPaymentController(paymentService)

	router := gin.Default()

	// set public route
	router.POST("/login", userController.Login)

	// set protected route (authenticated route)
	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.POST("/payment", paymentController.CreatePayment)
		protected.GET("/payments", paymentController.GetUserPayments)
		protected.POST("/logout", userController.Logout)
	}

	// start the server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
