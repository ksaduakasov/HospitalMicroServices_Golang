package main

import (
	"context"
	"flag"
	"github.com/Fring02/HospitalMicroservices/ReceptionService/pkg"
	"github.com/Fring02/HospitalMicroservices/ReceptionService/pkg/repositories"
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
	dsn := flag.String("dsn", "postgresql://localhost/hospital?user=postgres&password=Rubin1!!", "PostGreSQL")
	flag.Parse()
	var err error
	pkg.Conn, err = openDB(*dsn)
	if err != nil{
		log.Fatalf("Failed to connect to db: ", err)
	}
	orderRepository = repositories.NewOrderRepository(pkg.Conn)
	router  := gin.Default()
	RouteOrders(router)
	router.Run(":4000") // listen and serve on port 4000
}
