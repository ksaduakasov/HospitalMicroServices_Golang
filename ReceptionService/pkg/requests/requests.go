package requests

import (
	"context"
	"github.com/Fring02/HospitalMicroservices/ReceptionService/core"
	hospitalpb "github.com/Fring02/HospitalMicroservices/grpc"
	"io"
	"log"
)

func DepartmentByDiseaseId(diseaseId int, client hospitalpb.DepartmentServiceClient) *core.Department {
	req := &hospitalpb.DepartmentRequest{DiseaseId: int32(diseaseId)}
	ctx := context.Background()
	resp, err := client.GetDepartmentByDiseaseId(ctx, req)
	if err != nil {
		log.Printf("Failed to get department: %v", err)
		return nil
	}
	return &core.Department{
		Id:          int(resp.Id),
		Name:        resp.Name,
		Description: resp.Description,
		DiseaseId:   int(resp.DiseaseId),
	}
}
func getAvailableDoctorsFromDepartmentServer(dep *core.Department, client hospitalpb.DepartmentServiceClient) []*hospitalpb.DoctorsResponse {
	req := &hospitalpb.DoctorsRequest{DepartmentId: int32(dep.Id), Status: true, DiseaseId: int32(dep.DiseaseId)}
	ctx := context.Background()
	stream, err := client.GetDoctors(ctx, req)
	if err != nil {
		log.Fatalf("error while getting doctors stream %v", err)
	}
	defer stream.CloseSend()
	doctors := []*hospitalpb.DoctorsResponse{}
LOOP:
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break LOOP
			}
			log.Fatalf("error with response from department server: %v", err)
		}
		doctors = append(doctors, resp)
	}
	return doctors
}

func GetAvailableDoctors(dep *core.Department) []*hospitalpb.DoctorsResponse {
	return getAvailableDoctorsFromDepartmentServer(dep, hospitalpb.GrpcClient)
}
func GetDepartmentByDiseaseId(diseaseId int) *core.Department {
	return DepartmentByDiseaseId(diseaseId, hospitalpb.GrpcClient)
}
