package handlers

import (
	"net/http"
	"strconv"

	"github.com/MeNoln/golangapi/models"

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

//PostTodo ...
func PostTodo(c *gin.Context) {
	var data models.Todo

	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Bad data",
		})
	}

	err := providers.CreateTodo(&data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad request"})
	}

	c.Status(200)
}

//PutTodo ...
func PutTodo(c *gin.Context) {
	var data models.Todo
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "id must be a number"})
	}

	if c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad data",
		})
	}

	if ID != data.ID {
		c.JSON(400, gin.H{
			"message": "Todos doesnt match",
			"routeId": ID,
			"dataId":  data.ID,
		})
	}

	err = providers.UpdateTodo(&data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad request"})
	}

	c.Status(200)
}

//RemoveTodo ...
func RemoveTodo(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "id must be a number"})
	}

	err = providers.DeleteTodo(ID)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad request"})
	}

	c.Status(200)
}
