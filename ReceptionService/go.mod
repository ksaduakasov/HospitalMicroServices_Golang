module github.com/Fring02/HospitalMicroservices/ReceptionService

go 1.15

replace github.com/Fring02/HospitalMicroservices/grpc => ../grpc

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/jackc/pgx/v4 v4.10.1
	google.golang.org/grpc v1.36.0
)
