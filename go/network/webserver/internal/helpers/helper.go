package helpers

import (
    "encoding/json"
    "fmt"
    "os"

    "webserver/internal/models"
)

var TodosFile = "./todos.json"

func ReadTodos() ([]models.Todo, error) {
    file, err := os.ReadFile(TodosFile)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    var todos []models.Todo
    if err := json.Unmarshal(file, &todos); err != nil {
        fmt.Println(err)
        return nil, err
    }

    return todos, nil
}

func WriteTodos(todos []models.Todo) error {
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
