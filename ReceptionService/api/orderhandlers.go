package main

import (
	"github.com/Fring02/HospitalMicroservices/ReceptionService/core"
	"github.com/Fring02/HospitalMicroservices/ReceptionService/core/interfaces"
	"github.com/Fring02/HospitalMicroservices/ReceptionService/pkg/requests"
	"github.com/gin-gonic/gin"
	"strconv"
)
var jsonContentType = "application/json; charset=utf-8"
var orderRepository interfaces.IOrdersRepository

func RouteOrders(router *gin.Engine)  {
	router.GET("/orders", GetAllOrders)
	router.GET("/orders/:id", GetOrderById)
	router.POST("/orders", CreateOrder)
	router.DELETE("/orders/:id", DeleteOrder)
	router.PUT("/orders/:id", UpdateOrder)
	router.GET("/availableDoctors")
}

func GetAllOrders(c *gin.Context)  {
	orders := orderRepository.GetAllOrders()
	c.JSON(200, orders)
}

func GetOrderById(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	order := orderRepository.GetOrderById(id)
	c.JSON(200, order)
}

func CreateOrder(c *gin.Context)  {
	order := &core.Order{}
	err := c.BindJSON(order)
	if err != nil {
		c.Data(400, jsonContentType, []byte("Fill all fields"))
		return
	}
	if orderRepository.CreateOrder(*order) {
		c.Data(200, jsonContentType, []byte("Created order \n"))
		dep := requests.GetDepartmentByDiseaseId(order.DiseaseId)
		if dep == nil {
			c.Data(400, jsonContentType, []byte("Failed to find department by disease"))
		}
		availableDoctors := requests.GetAvailableDoctors(dep)
		if len(availableDoctors) == 0 {
			c.Data(200, jsonContentType, []byte("No available doctors for now. Wait"))
		} else {
			doctor := availableDoctors[0]
			c.JSON(200, doctor)
		}
	}
	c.Data(500, jsonContentType, []byte("Failed to create order"))
}

func DeleteOrder(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	order := orderRepository.GetOrderById(id)
	if order == nil {
		c.Data(400, jsonContentType, []byte("No such order with id"))
		return
	}
	if orderRepository.DeleteOrder(*order) {
		c.Data(200, jsonContentType, []byte("Deleted order"))
		return
	}
	c.Data(500, jsonContentType, []byte("Failed to delete order"))
}

func UpdateOrder(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	model := orderRepository.GetOrderById(id)
	order := &core.Order{}
	err = c.BindJSON(order)
	if err != nil {
		c.Data(400, jsonContentType, []byte("Fill all fields"))
		return
	}
	order.Id = id
	updateValues(model, order)
	if orderRepository.UpdateOrder(*order) {
		c.Data(200, jsonContentType, []byte("Updated order"))
		return
	}
	c.Data(500, jsonContentType, []byte("Failed to update order"))
}

func updateValues(order *core.Order, updateOrder *core.Order)  {
	if updateOrder.PatientId > 0 {
		order.PatientId = updateOrder.PatientId
	}
	if updateOrder.DiseaseId > 0 {
		order.PatientId = updateOrder.PatientId
	}
	if len(updateOrder.Title) > 0 {
		order.Title = updateOrder.Title
	}
	if len(updateOrder.Description) > 0 {
		order.Description = updateOrder.Description
	}
}
