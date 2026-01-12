package main

import "ecommapi/internal/database"

func init() {
	database.DBInit()
}

func main() {
	database.CreateUser("test4", "test4@test.com", "assadasdsd")
}
