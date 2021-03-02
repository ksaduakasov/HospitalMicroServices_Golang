module github.com/Fring02/HospitalMicroservices/DepartmentMicroservice

go 1.15

replace github.com/Fring02/HospitalMicroservices/grpc => ../grpc

require (
	github.com/Fring02/HospitalMicroservices/grpc v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/jackc/pgx/v4 v4.10.1
)
