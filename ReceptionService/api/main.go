package main

import (
	"context"
	"flag"
	"github.com/Fring02/HospitalMicroservices/ReceptionService/pkg"
	"github.com/Fring02/HospitalMicroservices/ReceptionService/pkg/repositories"
	hospitalpb "github.com/Fring02/HospitalMicroservices/grpc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log"
	"os"
)

func openDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Println("Connection for database couldn't be established")
		return nil, err
	}
	return pool, nil
}

//"postgresql://localhost/hospital?user=postgres&password=Rubin1!!"
func main() {
	dsn := flag.String("dsn", os.Getenv("CONN"), "PostGreSQL")
	flag.Parse()
	var err error
	pkg.Conn, err = openDB(*dsn)
	if err != nil {
		log.Fatalf("Failed to connect to db: ", err)
	}
	conn, err := grpc.Dial(os.Getenv("GRPC_SERVER_CONN"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	log.Println("Grpc client listening on ", os.Getenv("GRPC_SERVER_CONN"))
	hospitalpb.GrpcClient = hospitalpb.NewDepartmentServiceClient(conn)
	orderRepository = repositories.NewOrderRepository(pkg.Conn)
	router := gin.Default()
	RouteOrders(router)
	router.Run(os.Getenv("HOST")) // listen and serve on port 4000
}
