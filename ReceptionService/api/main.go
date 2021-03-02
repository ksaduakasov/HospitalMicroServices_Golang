package main

import (
	"context"
	"flag"
	"github.com/Fring02/HospitalMicroservices/ReceptionService/pkg"
	"github.com/jackc/pgx/pgxpool"
	"log"
 	"github.com/gin-gonic/gin"
)

func openDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Println("Connection for database is established")
		return nil, err
	}
	return pool, nil
}

func main() {
	dsn := flag.String("dsn", "postgresql://localhost/hospital?user=postgres&password=Rubin1!!", "PostGreSQL")
	flag.Parse()
	var err error
	pkg.Conn, err = openDB(*dsn)
	if err != nil{
		log.Fatalf("Failed to connect to db: ", err)
	}
	router  := gin.Default()
	router.GET("/example", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "example",
		}) // gin.H is a shortcut for map[string]interface{}
	})
	router.Run(":4000") // listen and serve on port 8080
}
