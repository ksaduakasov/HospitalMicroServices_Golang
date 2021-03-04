package main

import (
	"context"
	hospitalpb "github.com/Fring02/HospitalMicroservices/grpc"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0"+os.Getenv("HOST"))
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	hospitalpb.RegisterDepartmentServiceServer(s, &Server{})
	log.Println("Server is running on port ", os.Getenv("HOST"))
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}

type Server struct {
	hospitalpb.UnimplementedDepartmentServiceServer
}

func (s *Server) GetDepartmentByDiseaseId(ctx context.Context, req *hospitalpb.DepartmentRequest) (*hospitalpb.DepartmentResponse, error) {
	diseaseId := req.GetDiseaseId()
	var pool *pgxpool.Pool
	var err error
	pool, err = OpenDB(os.Getenv("CONN"))
	if err != nil {
		log.Printf("Failed to connect to db: ", err)
		return nil, err
	}
	dep := GetDepartmentByDiseaseId(pool, ctx, diseaseId)
	return &hospitalpb.DepartmentResponse{
		Id:          int32(dep.ID),
		Name:        dep.Name,
		Description: dep.Description,
		DiseaseId:   int32(dep.DiseaseID),
	}, nil
}
func GetDepartmentByDiseaseId(pool *pgxpool.Pool, ctx context.Context, diseaseId int32) *hospitalpb.Department {

	row := pool.QueryRow(ctx,
		"SELECT * FROM departments WHERE disease_id = $1", diseaseId)
	department := &hospitalpb.Department{}
	err := row.Scan(&department.ID, &department.Name, &department.Description, &department.DiseaseID)
	if err != nil {
		log.Printf("BAD GET REQUEST: %v", err)
		return nil
	}
	return department
}
func (s *Server) GetDoctors(req *hospitalpb.DoctorsRequest, stream hospitalpb.DepartmentService_GetDoctorsServer) error {
	departmentID := req.GetDepartmentId()
	isAvailable := req.GetStatus()
	diseaseID := req.GetDiseaseId()
	doctors, price := requestDispatcher(departmentID, diseaseID, isAvailable)
	for i := 0; i < len(doctors); i++ {
		response := &hospitalpb.DoctorsResponse{
			DoctorId:     int32(doctors[i].ID),
			FirstName:    doctors[i].Firstname,
			LastName:     doctors[i].Lastname,
			Patronymic:   doctors[i].Patronymic,
			DepartmentId: int32(doctors[i].DepartmentID),
			Price:        doctors[i].DoctorMultiplier * float64(price),
		}
		err := stream.Send(response)
		if err != nil {
			log.Printf("Error while sending %v", err)
		}
	}
	return nil
}
func OpenDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Println("Connection for database couldn't be established")
		return nil, err
	}
	return pool, nil
}
func requestDispatcher(departmentID int32, diseaseID int32, isAvailable bool) ([]*hospitalpb.Doctor, int32) {

	var pool *pgxpool.Pool
	var err error
	pool, err = OpenDB(os.Getenv("CONN"))
	if err != nil {
		log.Fatalf("Failed to connect to db: ", err)
	}
	var doctors []*hospitalpb.Doctor
	var price int32
	doctors, price = FindAvailableDoctors(pool, departmentID, diseaseID, isAvailable)
	return doctors, price
}

func FindAvailableDoctors(pool *pgxpool.Pool, departmentID int32, diseaseID int32, isAvailable bool) ([]*hospitalpb.Doctor, int32) {
	log.Println("Trying to get available doctors with depId: ", departmentID)
	rows, err := pool.Query(context.Background(),
		"SELECT * FROM doctors WHERE department_id = $1 AND isAvailable = $2", departmentID, isAvailable)
	if err != nil {
		log.Printf("BAD GET REQUEST: %v", err)
		log.Printf("Didn't found doctors with depId %v", departmentID)
		return nil, 0
	}
	defer rows.Close()
	var doctors []*hospitalpb.Doctor
	for rows.Next() {
		doctor := &hospitalpb.Doctor{}
		err = rows.Scan(&doctor.ID, &doctor.Firstname, &doctor.Lastname, &doctor.Patronymic, &doctor.Phone,
			&doctor.Email, &doctor.Description, &doctor.WorkExp, &doctor.DepartmentID, &doctor.DoctorMultiplier, &doctor.IsAvailable)
		if err != nil {
			return nil, 0
		}
		doctors = append(doctors, doctor)
	}

	row := pool.QueryRow(context.Background(),
		"SELECT price FROM diseases WHERE id = $1", diseaseID)
	var price int32
	err = row.Scan(&price)
	if err != nil {
		log.Printf("BAD GET REQUEST: %v", err)
		return nil, 0
	}
	return doctors, price

}
