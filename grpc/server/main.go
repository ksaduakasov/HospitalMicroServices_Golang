package server

import (
	"github.com/Alemkhan/HospitalMicroservices/DepartmentMicroservice/core"
	"github.com/Alemkhan/HospitalMicroservices/DepartmentMicroservice/database"
	"github.com/Alemkhan/HospitalMicroservices/DepartmentMicroservice/pkg/repositories"
	hospitalpb "github.com/Alemkhan/HospitalMicroservices/grpc"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type DepartmentService struct {
	hospitalpb.UnimplementedDepartmentServiceServer
}

func (d *DepartmentService) sendAvailableDoctors(req *hospitalpb.DoctorsRequest, stream hospitalpb.DepartmentService_GetDoctorsServer) error{
	departmentID := req.GetDepartmentId()
	isAvailable := req.GetStatus()
	diseaseID := req.GetDiseaseId()
	doctors, price:= requestDispatcher(departmentID, diseaseID, isAvailable)
	for i:=0; i < len(doctors); i++ {
		response := &hospitalpb.DoctorsResponse{
			DoctorId: int32(doctors[i].ID),
			FirstName: doctors[i].Firstname,
			LastName: doctors[i].Lastname,
			Patronymic: doctors[i].Patronymic,
			DepartmentId: int32(doctors[i].DepartmentID),
			Price: int32(doctors[i].DoctorMultiplier*float64(price)),
		}
		err := stream.Send(response)
		if err != nil {
			log.Printf("Error while sending %v", err)
		}
	}
	return nil
}

func requestDispatcher(departmentID int32, diseaseID int32, isAvailable bool) ([]*core.Doctor, int32){

	var pool *pgxpool.Pool
	var err error
	pool, err = database.OpenDB(os.Getenv("connection"))
	doctorRepository := repositories.NewDoctorRepository(pool)
	if err != nil{
		log.Fatalf("Failed to connect to db: ", err)
	}
	var doctors []*core.Doctor
	var price int32
	doctors, price = doctorRepository.FindAvailableDoctors(departmentID, diseaseID, isAvailable)
	return doctors, price

}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	hospitalpb.RegisterDepartmentServiceServer(s, &DepartmentService{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}