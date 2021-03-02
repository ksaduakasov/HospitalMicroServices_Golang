module github.com/Alemkhan/HospitalMicroservices/grpc

go 1.15

replace github.com/Alemkhan/HospitalMicroservices/DepartmentMicroservice => ../DepartmentMicroservice

require (
	github.com/Alemkhan/HospitalMicroservices/DepartmentMicroservice v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.2
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.25.0
)
