package handlers

import (
	"strconv"

	"github.com/MeNoln/golangapi/providers"
	"github.com/gin-gonic/gin"
)

//GetTodos ...
func GetTodos(c *gin.Context) {
	todos, err := providers.GetAllTodos()
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad request"})
	}
	c.JSON(200, todos)
}

//GetTodo ...
func GetTodo(c *gin.Context) {
	id := c.Param("id")
	todoID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Id must be a number"})
	}
	todo, err := providers.GetCurrentTodo(todoID)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad request"})
	}
	c.JSON(200, todo)
}
