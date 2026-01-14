package main

import (
	"ecommapi/docs"
	"ecommapi/internal/api/handlers"
	"ecommapi/internal/api/middlewares"
	db "ecommapi/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/gofor-little/env"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	db.DBInit()
}

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	router.GET("/health", handlers.Health)

	router.POST("/auth/register", handlers.Register)
	router.POST("/auth/login", handlers.Login)
	router.POST("/auth/logout", middlewares.CheckToken(), handlers.Logout)

	router.GET("/products", handlers.GetProducts)
	router.POST("/products", middlewares.CheckIsAdmin(), handlers.CreateProduct)

	router.GET("/inventory/:product_id", middlewares.CheckToken(), handlers.GetInventory)
	router.PATCH("/inventory/:product_id", middlewares.CheckIsAdmin(), handlers.UpdateInventory)

	// router.GET("/cart", middlewares.CheckIsCustomer(), handlers.GetCart)
	// router.POST("/cart/items", middlewares.CheckIsCustomer(), handlers.AddToCart)

	// router.POST("/orders", middlewares.CheckIsCustomer(), handlers.PlaceOrder)

	// router.GET("/notifications", middlewares.CheckIsCustomer(), handlers.GetNotifications)
	// router.PATCH("/notifications/:id/read", middlewares.CheckIsCustomer(), handlers.UpdateNotification)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	port, err := env.MustGet("PORT")
	if err != nil {
		panic(err)
	}
	router.Run(":" + port)
}
