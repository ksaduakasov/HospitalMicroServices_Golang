package main

import (
	"flag"
	"github.com/Fring02/HospitalMicroservices/DiseaseService/api"
	"github.com/Fring02/HospitalMicroservices/DiseaseService/database"
	"github.com/Fring02/HospitalMicroservices/DiseaseService/pkg/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

func main() {
	dsn := flag.String("dsn", os.Getenv("CONN"), "PostGreSQL")
	flag.Parse()
	var pool *pgxpool.Pool
	var err error
	pool, err = database.OpenDB(*dsn)
	api.DiseaseRepository = repositories.NewDiseaseRepository(pool)
	if err != nil {
		log.Fatalf("Failed to connect to db: ", err)
	}
	router := gin.Default()
	api.RouteDiseases(router)
	router.Run(":4003")
}
