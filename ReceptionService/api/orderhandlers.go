package main

import (
	"github.com/Fring02/HospitalMicroservices/ReceptionService/core"
	"github.com/Fring02/HospitalMicroservices/ReceptionService/pkg/repositories"
	"github.com/gin-gonic/gin"
	"strconv"
)
var jsonContentType = "application/json; charset=utf-8"
var orderRepository = repositories.NewOrderRepository()

func RouteOrders(router *gin.Engine)  {
	router.GET("/orders", GetAllOrders)
	router.GET("/orders/:id", GetOrderById)
	router.POST("/orders", CreateOrder)
	router.DELETE("/orders/:id", DeleteOrder)

}

func GetAllOrders(c *gin.Context)  {
	orders := orderRepository.GetAllOrders()
	c.JSON(200, orders)
}

func GetOrderById(c *gin.Context)  {
	id, err := strconv.Atoi(c.Request.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect id format"))
	}
	order := orderRepository.GetOrderById(id)
	c.JSON(200, order)
}

func CreateOrder(c *gin.Context)  {
	order := &core.Order{}
	err := c.BindJSON(order)
	if err != nil {
		c.Data(400, jsonContentType, []byte("Fill all fields"))
	}
	if orderRepository.CreateOrder(*order) {
		c.Data(200, jsonContentType, []byte("Created order"))
	}
	c.Data(500, jsonContentType, []byte("Failed to create order"))
}

func DeleteOrder(c *gin.Context)  {
	id, err := strconv.Atoi(c.Request.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect id format"))
	}
	order := orderRepository.GetOrderById(id)
	if order == nil {
		c.Data(400, jsonContentType, []byte("No such order with id"))
		return
	}
	if orderRepository.DeleteOrder(*order) {
		c.Data(200, jsonContentType, []byte("Deleted order"))
	}
	c.Data(500, jsonContentType, []byte("Failed to delete order"))
}

func UpdateOrder()  {
	
}