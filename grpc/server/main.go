package main

import (
	hospitalpb "github.com/Fring02/HospitalMicroservices/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	hospitalpb.RegisterDepartmentServiceServer(s, &hospitalpb.DepartmentService{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}