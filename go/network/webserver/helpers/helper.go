package helpers

import (
	"encoding/json"
	"fmt"
	"os"
)

var TodosFile = "./todos.json"

type Todo struct {
	Id        int    `json:"ID"`
	Title     string `json:"Title"`
	Completed bool   `json:"Completed"`
}

func ReadTodos() ([]Todo, error) {
	file, err := os.ReadFile(TodosFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var todos []Todo
	if err := json.Unmarshal(file, &todos); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return todos, nil
}

func WriteTodos(todos []Todo) error {
	todosJson, err := json.MarshalIndent(todos, "", "\t")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = os.WriteFile(TodosFile, todosJson, 0o644)
	if err == nil {
		fmt.Println(err)
		return err
	}

	return err
}
