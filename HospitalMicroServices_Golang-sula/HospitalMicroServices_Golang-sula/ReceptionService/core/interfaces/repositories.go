package interfaces

import "github.com/Fring02/HospitalMicroservices/ReceptionService/core"

type IOrdersRepository interface {
	CreateOrder(order core.Order) bool
	GetAllOrders() []*core.Order
	GetOrderById(id int) *core.Order
	DeleteOrder(order core.Order) bool
	UpdateOrder(order core.Order) bool
}
