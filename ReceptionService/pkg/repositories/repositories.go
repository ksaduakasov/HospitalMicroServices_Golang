package repositories

import (
	"context"
	"github.com/Fring02/HospitalMicroservices/ReceptionService/core"
	"github.com/Fring02/HospitalMicroservices/ReceptionService/core/interfaces"
	"github.com/Fring02/HospitalMicroservices/ReceptionService/pkg"
	"github.com/jackc/pgx/pgxpool"
	"log"
)
type OrderRepository struct {
	pool pgxpool.Pool
}

func NewOrderRepository() interfaces.IOrdersRepository {
	return &OrderRepository{pool: *pkg.Conn}
}
func (r *OrderRepository) CreateOrder(order core.Order) bool {
	sql := "INSERT INTO patient_orders(id, disease_id, patient_id, title, description) " +
		"VALUES($1, $2, $3, $4, $5) RETURNING id";
	row := r.pool.QueryRow(context.Background(),
		sql, order.Id, order.DiseaseId, order.PatientId, order.Title, order.Description)
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Printf("Unable to INSERT: %v\n", err)
		return false
	}
	return true
}

func (r OrderRepository) GetAllOrders() []*core.Order {
	stmt := "SELECT * FROM patient_orders"
	rows, err := r.pool.Query(context.Background(), stmt)
	if err != nil {
		log.Fatal("Failed to SELECT: %v", err)
		return nil
	}
	defer rows.Close()
	orders := []*core.Order{}
	for rows.Next() {
		o := &core.Order{}
		err = rows.Scan(&o.Id, &o.DiseaseId, &o.PatientId, &o.Title, &o.Description)
		if err != nil {
			log.Fatalf("Failed to scan: %v", err)
			return nil
		}
		orders = append(orders, o)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return orders
}
func (r* OrderRepository) GetOrderById(id int) *core.Order {
	stmt := "SELECT * FROM patient_orders WHERE id = $1"
	o := &core.Order{}
	err := r.pool.QueryRow(context.Background(), stmt, id).Scan(&o.Id, &o.DiseaseId, &o.PatientId, &o.Title, &o.Description)
	if err != nil {
		log.Println("Didn't find order with id ", id)
		return nil
	}
	return o
}

func (r *OrderRepository) DeleteOrder(order core.Order) bool  {
	_, err := r.pool.Exec(context.Background(),
		"DELETE FROM patient_orders WHERE id = $1", order.Id)
	if err != nil {
		return false
	}
	return true
}
func (r OrderRepository) UpdateOrder(order core.Order) bool {
	_, err := r.pool.Exec(context.Background(),
		"UPDATE patient_orders SET disease_id = $1, patient_id = $2, title = $3, description = $4 WHERE id = $5",
		order.DiseaseId, order.PatientId, order.Title, order.Description, order.Id)
	if err != nil {
		return false
	}
	return true
}