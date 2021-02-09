package main
import (
	"context"
	"flag"
	_ "fmt"
	_ "github.com/bmizerany/pat"
	"github.com/golangcollege/sessions"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/justinas/alice"
	"log"
	"net/http"
	"os"
	"time"
)
var App = &Application{
	ErrorLog: *log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	InfoLog:  *log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
}
func openDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Println("Connection for database is established")
		return nil, err
	}
	return pool, nil
}
var conn *pgxpool.Pool

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret")
	dsn := flag.String("dsn", "postgresql://localhost/hospital?user=postgres&password=Rubin1!!", "PostGreSQL")
	flag.Parse()
	var err error
	conn, err = openDB(*dsn)
	if err != nil {
		App.ErrorLog.Fatalf("Database connection failed: %v\n", err)
	}
	defer conn.Close()
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	App.Session = session
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: &App.ErrorLog,
		Handler:  App.Routes(),
	}
	App.InfoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	App.ErrorLog.Fatal(err)
}