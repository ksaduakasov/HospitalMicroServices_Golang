package repositories

import (
	"context"
	"github.com/Fring02/HospitalMicroservices/UserService/core"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type UserRepository struct {
	conn *pgxpool.Pool
}

func NewUserRepository(connection *pgxpool.Pool) UserRepository {
	return UserRepository{conn: connection}
}


func (d *UserRepository) GetUsers() []*core.User  {

	rows, err := d.conn.Query(context.Background(),
		"SELECT * FROM patients")
	if err != nil {
		log.Printf("BAD GET REQUEST: %v", err)
		return nil
	}
	defer rows.Close()
	var users []*core.User
	for rows.Next() {
		user := &core.User{}
		err = rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Balance, &user.Patronymic, &user.Email, &user.Password)
		if err != nil {
			return nil
		}
		users = append(users, user)
	}
	return users

}

func (d *UserRepository) GetUserByID(id int) *core.User  {

	row := d.conn.QueryRow(context.Background(),
		"SELECT * FROM patients WHERE id = $1", id)
	user := &core.User{}
	err := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Balance, &user.Patronymic, &user.Email, &user.Password)
	if err != nil {
		log.Printf("BAD GET REQUEST: %v", err)
		return nil
	}
	return user

}

func (d *UserRepository) CreateUser(user core.User) (int, error) {

	row := d.conn.QueryRow(context.Background(),
		"INSERT INTO patients (firstname, lastname, balance, patronymic, email, password)" +
			"VALUES($1,$2,$3,$4,$5,$6) RETURNING id",
		user.Firstname, user.Lastname, user.Balance, user.Patronymic,
		user.Email, user.Password)
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Printf("Creation fail %v", err)
		return 0, err
	}
	return id, nil

}

func (d *UserRepository) DeleteUser(id int) (bool, error)  {

	_, err := d.conn.Exec(context.Background(),
		"DELETE FROM patients WHERE id = $1", id)
	if err != nil {
		log.Printf("BAD DELETE REQUEST")
		return false, err
	}
	return true, nil

}

func (d *UserRepository) UpdateUser(doctor core.User)  (bool, error){

	_, err := d.conn.Exec(context.Background(),
		"UPDATE patients SET firstname = $1, lastname = $2, balance = $3, patronymic = $4, email = $5, password = $6 WHERE id = $7",
		doctor.Firstname, doctor.Lastname, doctor.Balance, doctor.Patronymic, doctor.Email, doctor.Password, doctor.ID)
	if err != nil {
		log.Printf("BAD PUT REQUEST")
		return false, err
	}
	return true, nil

}

func (d *UserRepository) GetUser(email, password string)  *core.User {

	row := d.conn.QueryRow(context.Background(),
		"SELECT * FROM patients WHERE email = $1 AND password = $2", email, password)
	user := &core.User{}
	err := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Balance, &user.Patronymic, &user.Email, &user.Password)
	if err != nil {
		log.Printf("BAD GET REQUEST: %v", err)
		return nil
	}
	return user

}


