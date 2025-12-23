package handlers

import (
	"net/http"
	"strconv"

	"webserver/helpers"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// GetTodos godoc
// @Summary Get ToDos
// @Description Retrieve all ToDos
// @Tags todos
// @Produce json
// @Success 200 {object} []helpers.Todo "Details of ToDos"
// @Router /todos [get]
func GetTodos(c *gin.Context) {
	todos, err := helpers.ReadTodos()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {

}

// AddTodo godoc
// @Summary Add ToDo
// @Description Creating a new ToDo
// @Tags todos
// @Accept json
// @Produce json
// @Param data body helpers.Todo true "New ToDo"
// @Success 200 {object} helpers.Todo "Details of the ToDo"
// @Router /todos [post]
func AddTodo(c *gin.Context) {
	var todo helpers.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	todos, err := helpers.ReadTodos()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	todos = append(todos, todo)
	err = helpers.WriteTodos(todos)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(200, todo)
}

func UpdateTodo(c *gin.Context) {

}

// DeleteTodo godoc
// @Summary Delete ToDo
// @Description Deleting the ToDo
// @Tags todos
// @Produce json
// @Param id path int true "ToDo ID"
// @Success 200 {object} helpers.Todo "Details of the ToDo"
// @Router /todos/:id [delete]
func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	todos, err := helpers.ReadTodos()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	for i, td := range todos {
		if td.Id == id {
			todos = append(todos[:i], todos[i+1:]...)

			err = helpers.WriteTodos(todos)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}

			c.JSON(200, td)
			return
		}
	}

	c.AbortWithStatus(http.StatusBadRequest)
}
