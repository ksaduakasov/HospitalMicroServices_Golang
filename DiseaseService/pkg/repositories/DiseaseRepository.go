package repositories

import (
	"context"
	"github.com/Fring02/HospitalMicroservices/DiseaseService/core"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type DiseaseRepository struct {
	conn *pgxpool.Pool
}

func NewDiseaseRepository(connection *pgxpool.Pool) DiseaseRepository {
	return DiseaseRepository{conn: connection}
}

func (d *DiseaseRepository) CreateDisease(disease core.Disease) (int, error) {

	row := d.conn.QueryRow(context.Background(),
		"INSERT INTO diseases (bodypart, price)"+
			"VALUES($1, $2) RETURNING id",
		disease.BodyPart, disease.Price)
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Printf("Creation fail %v", err)
		return 0, err
	}
	return id, nil

}

func (d *DiseaseRepository) GetDiseases() []*core.Disease {

	rows, err := d.conn.Query(context.Background(),
		"SELECT * FROM diseases")
	if err != nil {
		log.Printf("BAD GET REQUEST: %v", err)
		return nil
	}
	defer rows.Close()
	var diseases []*core.Disease
	for rows.Next() {
		disease := &core.Disease{}
		err = rows.Scan(&disease.ID, &disease.BodyPart, &disease.Price)
		if err != nil {
			return nil
		}
		diseases = append(diseases, disease)
	}
	return diseases

}

func (d *DiseaseRepository) GetDiseaseByID(id int) *core.Disease {

	row := d.conn.QueryRow(context.Background(),
		"SELECT * FROM diseases WHERE id = $1", id)
	disease := &core.Disease{}
	err := row.Scan(&disease.ID, &disease.BodyPart, &disease.Price)
	if err != nil {
		log.Printf("BAD GET REQUEST: %v", err)
		return nil
	}
	return disease

}

func (d *DiseaseRepository) DeleteDisease(id int) (bool, error) {

	_, err := d.conn.Exec(context.Background(),
		"DELETE FROM diseases WHERE id = $1", id)
	if err != nil {
		log.Printf("BAD DELETE REQUEST")
		return false, err
	}
	return true, nil

}

func (d *DiseaseRepository) UpdateDisease(disease core.Disease) (bool, error) {

	_, err := d.conn.Exec(context.Background(),
		"UPDATE diseases SET bodypart = $1, price = $2 WHERE id = $3",
		disease.BodyPart, disease.Price, disease.ID)
	if err != nil {
		log.Printf("BAD PUT REQUEST")
		return false, err
	}
	return true, nil

}

func (d *DiseaseRepository) CheckForDisease(disease core.Disease) bool {
	row := d.conn.QueryRow(context.Background(),
		"SELECT * FROM diseases WHERE bodypart = $1", disease.BodyPart)
	err := row.Scan(&disease.ID, &disease.BodyPart)
	if err != nil {
		log.Printf("Check failed: %v\n", err)
		return false
	}
	return disease.ID > 0
}
