package main

import (
	"webserver/internal/http/handlers"
)

func main() {
	router := handlers.SetupRouter()

	err := router.Run()
	if err != nil {
		return
	}
}
