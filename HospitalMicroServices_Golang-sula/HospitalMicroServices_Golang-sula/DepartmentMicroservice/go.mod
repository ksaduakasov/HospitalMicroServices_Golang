module github.com/Fring02/HospitalMicroservices/DepartmentMicroservice

go 1.15

replace github.com/Fring02/HospitalMicroservices/grpc => ../grpc

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jackc/pgx/v4 v4.10.1
	google.golang.org/protobuf v1.25.0 // indirect
)
