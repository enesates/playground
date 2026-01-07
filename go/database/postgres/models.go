package main

import "time"

type User struct {
	id       int
	name     string
	email    string
	createAt time.Time
}
