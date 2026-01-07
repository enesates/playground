package main

import (
	"log"
)

func CreateTables() {
	err := GormDB.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	err = GormDB.AutoMigrate(&Profile{})
	if err != nil {
		log.Fatal(err)
	}
}

func AddUser(user User) {
	GormDB.Create(&User{Name: user.Name, Email: user.Email, Age: user.Age})
}

func GetUsers() []User {
	var users []User
	GormDB.Find(&users)
	return users
}
