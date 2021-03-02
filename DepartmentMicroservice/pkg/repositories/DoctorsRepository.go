package repositories

import (
	"github.com/Alemkhan/HospitalMicroservices/DepartmentMicroservice/core"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"context"
)

type DoctorRepository struct {
	conn *pgxpool.Pool
}

func NewDoctorRepository(connection *pgxpool.Pool) DoctorRepository {
	return DoctorRepository{conn: connection}
}

func (d *DoctorRepository) CreateDoctor(doctor core.Doctor) (int, error) {

	row := d.conn.QueryRow(context.Background(),
		"INSERT INTO doctors(firstname, lastname, patronymic, phone, email, description, workexp, department_id,isAvailable, doctorMultiplier)" +
			"VALUES($1,$2,$3,$4,$5,$6,$7,$8, true,$9) RETURNING id",
		doctor.Firstname, doctor.Lastname, doctor.Patronymic, doctor.Phone,
		doctor.Email, doctor.Description, doctor.WorkExp, doctor.DepartmentID, float64(1.0 + (float64(doctor.WorkExp) / 10.0)))
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Printf("Creation fail %v", err)
		return 0, err
	}
	return id, nil

}

func (d *DoctorRepository) GetDoctors() []*core.Doctor  {

	rows, err := d.conn.Query(context.Background(),
		"SELECT * FROM doctors")
	if err != nil {
		log.Printf("BAD GET REQUEST: %v", err)
		return nil
	}
	defer rows.Close()
	var doctors []*core.Doctor
	for rows.Next() {
		doctor := &core.Doctor{}
		err = rows.Scan(&doctor.ID, &doctor.Firstname, &doctor.Lastname, &doctor.Patronymic, &doctor.Phone,
			&doctor.Email, &doctor.Description, &doctor.WorkExp, &doctor.DepartmentID, &doctor.IsAvailable,
			&doctor.DoctorMultiplier)
		if err != nil {
			return nil
		}
		doctors = append(doctors, doctor)
	}
	return doctors

}

func (d *DoctorRepository) GetDoctorByID(id int) *core.Doctor  {

	row := d.conn.QueryRow(context.Background(),
		"SELECT * FROM doctors WHERE id = $1", id)
	doctor := &core.Doctor{}
	err := row.Scan(&doctor.ID, &doctor.Firstname, &doctor.Lastname, &doctor.Patronymic, &doctor.Phone,
		&doctor.Email, &doctor.Description, &doctor.WorkExp, &doctor.DepartmentID, &doctor.IsAvailable,
		&doctor.DoctorMultiplier)
	if err != nil {
		log.Printf("BAD GET REQUEST: %v", err)
		return nil
	}
	return doctor

}

func (d *DoctorRepository) DeleteDoctor(id int) (bool, error)  {

	_, err := d.conn.Exec(context.Background(),
		"DELETE FROM doctors WHERE id = $1", id)
	if err != nil {
		log.Printf("BAD DELETE REQUEST")
		return false, err
	}
	return true, nil

}

func (d *DoctorRepository) UpdateDoctor(doctor core.Doctor)  (bool, error){

	_, err := d.conn.Exec(context.Background(),
		"UPDATE doctors SET firstname = $1, lastname = $2, patronymic = $3, phone = $4, email = $5, description = $6, " +
			"workexp = $7, department_id = $8, isAvailable = $9, doctorMultiplier = $10",
			doctor.Firstname, doctor.Lastname, doctor.Patronymic, doctor.Phone,
			doctor.Email, doctor.Description, doctor.WorkExp, doctor.DepartmentID, doctor.IsAvailable, float64(1.0 + (float64(doctor.WorkExp) / 10.0)))
	if err != nil {
		log.Printf("BAD PUT REQUEST")
		return false, err
	}
	return true, nil

}