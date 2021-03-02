module github.com/Fring02/HospitalMicroservices/ReceptionService

go 1.15

replace github.com/Fring02/HospitalMicroservices/grpc v0.0.0-20210302125455-08b7d77f09c7 => ../grpc

require (
	github.com/Fring02/HospitalMicroservices/grpc v0.0.0-20210302125455-08b7d77f09c7
	github.com/gin-gonic/gin v1.6.3
	github.com/jackc/pgx/v4 v4.10.1
	google.golang.org/grpc v1.36.0
)
