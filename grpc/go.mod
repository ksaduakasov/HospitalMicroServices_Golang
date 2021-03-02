module github.com/Fring02/HospitalMicroservices/grpc

go 1.15

replace github.com/Fring02/HospitalMicroservices/DepartmentMicroservice => ../DepartmentMicroservice

require (
	github.com/Fring02/HospitalMicroservices/DepartmentMicroservice v0.0.0-00010101000000-000000000000
	github.com/jackc/pgx/v4 v4.10.1
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.25.0
)
