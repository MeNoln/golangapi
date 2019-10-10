package main

import (
	"log"

	"github.com/MeNoln/golangapi/db"
	Handler "github.com/MeNoln/golangapi/handlers"
	"github.com/gin-gonic/gin"
)

func todoRoutes(router *gin.Engine) {
	router.GET("api/v1/todos", Handler.GetTodos)
	router.GET("api/v1/todo/:id", Handler.GetTodo)
	router.POST("api/v1/todo", Handler.PostTodo)
	router.PUT("api/v1/todo/:id", Handler.PutTodo)
	router.DELETE("api/v1/todo/:id", Handler.RemoveTodo)
}

func orderRoutes(router *gin.Engine) {
	router.GET("api/v2/orders", Handler.GetOrders)
	router.GET("api/v2/order/:id", Handler.GetOrder)
	router.POST("api/v2/order", Handler.PostOrder)
	router.PUT("api/v2/order/:id", Handler.PutOrder)
	router.DELETE("api/v2/order/:id", Handler.RemoveOrder)
}

func main() {
	log.Println("Starting server")

	db.InitializeDB()

	router := gin.Default()

	todoRoutes(router)
	orderRoutes(router)

	router.Run(":7000")
}
