package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Fring02/HospitalMicroservices/ReceptionService/core"
	"io"
	"log"
	"net/http"
)

func GetDepartmentByDiseaseId(diseaseId int) *core.Department {
	url := fmt.Sprintf("http://localhost:4001/departments/%v",diseaseId)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to get department: %v", err)
		return nil
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	dep := &core.Department{}
	decoder.Decode(dep)
	return dep
}
func getAvailableDoctorsFromDepartmentServer(dep *core.Department, client hospitalpb.DepartmentServiceClient) []*hospitalpb.DoctorsResponse {
	req := &hospitalpb.DoctorsRequest{DepartmentId: int32(dep.Id), Status: true}
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
			if err == io.EOF{
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
