package repositories

import (
	"context"
	"github.com/Fring02/HospitalMicroservices/DepartmentMicroservice/core"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type DepartmentsRepository struct {
	conn *pgxpool.Pool
}

func NewDepartmentsRepository(connection *pgxpool.Pool) DepartmentsRepository {
	return DepartmentsRepository{conn: connection}
}

func (d *DepartmentsRepository) CreateDepartment(department core.Department) (int, error) {

	row := d.conn.QueryRow(context.Background(),
		"INSERT INTO departments(name, description, disease_id)"+
			"VALUES($1,$2,$3) RETURNING id",
		department.Name, department.Description, department.DiseaseID)
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Printf("Creation fail %v", err)
		return 0, err
	}
	return id, nil

}

func (d *DepartmentsRepository) GetDepartments() []*core.Department {

	rows, err := d.conn.Query(context.Background(),
		"SELECT * FROM departments")
	if err != nil {
		log.Printf("BAD GET REQUEST: %v", err)
		return nil
	}
	defer rows.Close()
	var departments []*core.Department
	for rows.Next() {
		department := &core.Department{}
		err = rows.Scan(&department.ID, &department.Name, &department.Description, &department.DiseaseID)
		if err != nil {
			return nil
		}
		departments = append(departments, department)
	}
	return departments
}

func (d *DepartmentsRepository) GetDepartmentsByDiseaseId(diseaseId int) []*core.Department {

	rows, err := d.conn.Query(context.Background(),
		"SELECT * FROM departments WHERE disease_id = $1", diseaseId)
	if err != nil {
		log.Printf("BAD GET REQUEST: %v", err)
		return nil
	}
	defer rows.Close()
	var departments []*core.Department
	for rows.Next() {
		department := &core.Department{}
		err = rows.Scan(&department.ID, &department.Name, &department.Description, &department.DiseaseID)
		if err != nil {
			return nil
		}
		departments = append(departments, department)
	}
	return departments
}

func (d *DepartmentsRepository) GetDepartmentByID(id int) *core.Department {

	row := d.conn.QueryRow(context.Background(),
		"SELECT * FROM departments WHERE id = $1", id)
	department := &core.Department{}
	err := row.Scan(&department.ID, &department.Name, &department.Description, &department.DiseaseID)
	if err != nil {
		log.Printf("BAD GET REQUEST: %v", err)
		return nil
	}
	return department

}

func (d *DepartmentsRepository) DeleteDepartment(id int) (bool, error) {

	_, err := d.conn.Exec(context.Background(),
		"DELETE FROM departments WHERE id = $1", id)
	if err != nil {
		log.Printf("BAD DELETE REQUEST")
		return false, err
	}
	return true, nil

}

func (d *DepartmentsRepository) UpdateDepartment(department core.Department) (bool, error) {

	_, err := d.conn.Exec(context.Background(),
		"UPDATE departments SET name = $1, description = $2, disease_id = $3",
		department.Name, department.Description, department.DiseaseID)
	if err != nil {
		log.Printf("BAD PUT REQUEST")
		return false, err
	}
	return true, nil

}
