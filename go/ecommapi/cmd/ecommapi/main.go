package main

import (
	"ecommapi/api/handlers"
	"ecommapi/docs"
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
	router.POST("/auth/register", handlers.RegisterUser)
	router.POST("/auth/login", handlers.LoginUser)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	port, err := env.MustGet("PORT")
	if err != nil {
		panic(err)
	}
	router.Run(":" + port)
}
