package main

import (
	"ecommapi/docs"
	"ecommapi/internal/api"
	"ecommapi/internal/cart"
	db "ecommapi/internal/helpers/database"
	"ecommapi/internal/inventory"
	"ecommapi/internal/notification"
	"ecommapi/internal/order"
	"ecommapi/internal/product"
	"ecommapi/internal/user"

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

	router.GET("/health", api.Health)

	router.POST("/auth/register", user.Register)
	router.POST("/auth/login", user.Login)
	router.POST("/auth/logout", api.CheckToken(), user.Logout)

	router.GET("/products", product.GetProducts)
	router.POST("/products", api.CheckIsAdmin(), product.AddProduct)

	router.GET("/inventory/:product_id", api.CheckToken(), inventory.GetInventory)
	router.PATCH("/inventory/:product_id", api.CheckIsAdmin(), inventory.CreateOrUpdateInventory)

	router.GET("/cart", api.CheckIsCustomer(), cart.GetCart)
	router.POST("/cart/items", api.CheckIsCustomer(), cart.AddToCart)

	router.POST("/orders", api.CheckIsCustomer(), order.PlaceOrder)

	router.GET("/notifications", api.CheckIsCustomer(), notification.GetNotifications)
	router.PATCH("/notifications/:id/read", api.CheckIsCustomer(), notification.MakeNotificationRead)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	port, err := env.MustGet("PORT")
	if err != nil {
		panic(err)
	}

	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}
