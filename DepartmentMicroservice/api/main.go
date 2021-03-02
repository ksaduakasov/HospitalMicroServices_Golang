package main

import (
	"context"
	"flag"
	"github.com/Alemkhan/HospitalMicroservices/DepartmentMicroservice/pkg/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func openDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Println("Connection for database is established")
		return nil, err
	}
	return pool, nil
}

func main() {
	dsn := flag.String("dsn", "postgresql://localhost/hospital?user=postgres&password=1337", "PostGreSQL")
	flag.Parse()
	var pool *pgxpool.Pool
	var err error
	pool, err = openDB(*dsn)
	doctorRepository = repositories.NewDoctorRepository(pool)
	if err != nil{
		log.Fatalf("Failed to connect to db: ", err)
	}
	router  := gin.Default()
	RouteDoctors(router)
	router.Run(":4000")
}
