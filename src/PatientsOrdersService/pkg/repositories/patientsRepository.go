package repositories

import (
	"context"
	"github.com/Fring02/HospitalMicroservices/src/PatientsOrdersService/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func CreatePatient(conn *pgxpool.Pool, model models.Patient) (int, error) {
	row := conn.QueryRow(context.Background(),
		"INSERT INTO patients(firstname, lastname, patronymic, balance) " +
			"VALUES($1,$2,$3,$4) RETURNING id",
		model.Firstname, model.Lastname, model.Patronymic, model.Balance)
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Fatalln("Failed to create")
		return 0, err
	}
	return id, nil
}
func GetPatientById(conn *pgxpool.Pool, id int) (*models.Patient, error) {
	row := conn.QueryRow(context.Background(),
		"SELECT * FROM patients WHERE id = $1",id)
	patient := &models.Patient{}
	err := row.Scan(&patient.Id, &patient.Firstname, &patient.Lastname, &patient.Balance, &patient.Patronymic)
	if err != nil {
		log.Fatalln("Failed to create patient: ", err)
		return nil, err
	}
	return patient, nil
}
func GetAllPatients(conn *pgxpool.Pool) ([]*models.Patient, error) {
	rows, err := conn.Query(context.Background(),
		"SELECT * FROM patients")
	if err != nil {
		log.Fatal("Failed to get all patients: ", err)
		return nil, err
	}
	defer rows.Close()
	var patients []*models.Patient
	for rows.Next(){
		patient := &models.Patient{}
		err = rows.Scan(&patient.Id, &patient.Firstname, &patient.Lastname, &patient.Balance, &patient.Patronymic)
		if err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}
	return patients, nil
}

func DeletePatient(conn *pgxpool.Pool, id int) (bool, error) {
	_, err := conn.Exec(context.Background(),
		"DELETE FROM patients WHERE id = $1",id)
	if err != nil {
		return false, err
	}
	return true, nil
}
func UpdatePatient(conn *pgxpool.Pool, model models.Patient, oldId int) (bool, error) {
	_, err := conn.Exec(context.Background(),
		"UPDATE patients SET id = $1, firstname = $2, lastname = $3, balance = $4, patronymic = $5 WHERE id = $6",
		model.Id, model.Firstname, model.Lastname, model.Balance, model.Patronymic, oldId)
	if err != nil {
		return false, err
	}
	return true, nil
}