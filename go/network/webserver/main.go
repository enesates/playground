package main

import (
	"webserver/docs"
	"webserver/handlers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /
// @title   CRUD API
// @version  1.0
// @description A CRUD Backend API

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	router.GET("/health", handlers.Health)
	router.GET("/todos", handlers.GetTodos)
	router.GET("/todos/:id", handlers.GetTodo)
	router.POST("/todos", handlers.AddTodo)
	router.PUT("/todos/:id", handlers.UpdateTodo)
	router.DELETE("/todos/:id", handlers.DeleteTodo)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run()
}
