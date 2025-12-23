package main

import (
	_ "gingonic/docs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type UserRequest struct {
	Msg string `json:"msg"`
}

// Get Time
// @Summary get time example
// @Schemes
// @Description get time
// @Success 200
// @Router /time [get]
func getTime(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": time.Now(),
	})
}

// PostExample
// @Summary post example
// @Schemes
// @Description post example
// @Accept json
// @Produce json
// @Success 200
// @Router /echo [post]
func postEcho(c *gin.Context) {
	var ur UserRequest

	if err := c.ShouldBindJSON(&ur); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(200, ur)
}

// @title TimeEcho API
// @version 1.0
// @description Swagger Example.
// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()

	router.GET("/time", getTime)

	router.POST("/echo", postEcho)

	//http.Handle("/swagger/", httpSwagger.WrapHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run() // listens on 0.0.0.0:8080 by default
}
