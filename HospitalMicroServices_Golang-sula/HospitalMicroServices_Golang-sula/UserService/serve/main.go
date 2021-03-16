package main

import (
	"flag"
	"github.com/Fring02/HospitalMicroservices/UserService/api"
	"github.com/Fring02/HospitalMicroservices/UserService/database"
	"github.com/Fring02/HospitalMicroservices/UserService/pkg/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func main() {
	dsn := flag.String("dsn", "postgresql://localhost/hospital?user=postgres&password=Kalbek2002", "PostGreSQL")
	flag.Parse()
	var pool *pgxpool.Pool
	var err error
	pool, err = database.OpenDB(*dsn)
	api.UserRepository = repositories.NewUserRepository(pool)
	if err != nil {
		log.Fatalf("Failed to connect to db: ", err)
	}
	router := gin.Default()
	api.RouteUsers(router)
	router.Run(":4001")
}
