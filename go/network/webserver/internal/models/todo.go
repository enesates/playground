package models

type Todo struct {
	Id        int    `json:"ID"`
	Title     string `json:"Title"`
	Completed bool   `json:"Completed"`
}
