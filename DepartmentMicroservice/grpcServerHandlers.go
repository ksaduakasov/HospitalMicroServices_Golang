package DepartmentMicroservice

import (
	"github.com/Fring02/HospitalMicroservices/DepartmentMicroservice/core"
	"github.com/Fring02/HospitalMicroservices/DepartmentMicroservice/database"
	"github.com/Fring02/HospitalMicroservices/DepartmentMicroservice/pkg/repositories"
	hospitalpb "github.com/Fring02/HospitalMicroservices/grpc"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

func sendAvailableDoctors(req *hospitalpb.DoctorsRequest, stream hospitalpb.DepartmentService_GetDoctorsServer) error{
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