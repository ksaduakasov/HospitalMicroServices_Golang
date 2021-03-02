package main

import (
	hospitalpb "github.com/Fring02/HospitalMicroservices/grpc"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	log.Printf("Grpc client listening on 50051")
	hospitalpb.GrpcClient = hospitalpb.NewDepartmentServiceClient(conn)
}
