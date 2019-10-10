package handlers

import (
	"log"
	"strconv"

	"github.com/MeNoln/golangapi/models"

	"github.com/MeNoln/golangapi/providers"
	"github.com/gin-gonic/gin"
)

//GetOrders ...
func GetOrders(c *gin.Context) {
	orders, err := providers.GetAllOrders()
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad request"})
	}
	c.JSON(200, orders)
}

//GetOrder ...
func GetOrder(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "id must be integer",
		})
	}

	order, err := providers.GetCurrentOrder(ID)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Bad request",
		})
	}
	c.JSON(200, order)
}

//PostOrder ...
func PostOrder(c *gin.Context) {
	var data models.OrderResponseModel

	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Bad request",
		})
	}
	log.Println(data.OrderAmount)
	err := providers.CreateNewOrder(&data)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Bad request",
		})
	}
	c.Status(200)
}

//PutOrder ...
func PutOrder(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "id must be a number"})
	}
	var data models.OrderResponseModel
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "Bad data"})
	}

	if ID != data.ID {
		c.AbortWithStatusJSON(400, gin.H{"message": "Id doesnt match"})
	}

	err = providers.UpdateOrder(&data)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "Bad data"})
	}
	c.Status(200)
}

//RemoveOrder ...
func RemoveOrder(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "id must be a number"})
	}

	err = providers.DeleteOrder(ID)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": "Bad data"})
	}
	c.Status(200)
}
