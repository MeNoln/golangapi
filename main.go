package main

import (
	"log"

	"github.com/MeNoln/golangapi/db"
	"github.com/MeNoln/golangapi/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server")

	db.InitializeDB()

	router := gin.Default()

	router.GET("/api/todos", handlers.GetTodos)
	router.GET("/api/todo/:id", handlers.GetTodo)

	router.Run(":7000")
}
