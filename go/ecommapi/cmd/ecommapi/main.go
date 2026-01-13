package main

import (
	// "ecommapi/docs"
	"ecommapi/internal/database"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func init() {
	database.DBInit()
}

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	// router.GET("/health", handlers.Health)
	// router.GET("/auth/register", handlers.GetTodos)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run()
}
