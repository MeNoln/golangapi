package main

import (
	"log"

	"github.com/MeNoln/golangapi/db"
	TodoHandler "github.com/MeNoln/golangapi/handlers"
	"github.com/gin-gonic/gin"
)

func todoRoutes(router *gin.Engine) {
	router.GET("api/v1/todos", TodoHandler.GetTodos)
	router.GET("api/v1/todo/:id", TodoHandler.GetTodo)
	router.POST("api/v1/todo", TodoHandler.PostTodo)
	router.PUT("api/v1/todo/:id", TodoHandler.PutTodo)
	router.DELETE("api/v1/todo/:id", TodoHandler.RemoveTodo)
}

func main() {
	log.Println("Starting server")

	db.InitializeDB()

	router := gin.New()

	todoRoutes(router)

	router.Run(":7000")
}
