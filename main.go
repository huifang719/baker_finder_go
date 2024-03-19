package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/huifang719/baker_finder_go/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type config struct {
	port int
	env string
	DB *database.Queries
}

type application struct {
	config config
	infoLog *log.Logger
	errorLog *log.Logger
	version string
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
		IdleTimeout: 30*time.Second,
		ReadTimeout: 5*time.Second,
		WriteTimeout: 5*time.Second,
	}
		
	app.infoLog.Printf("Starting %s server on %d", app.config.env, app.config.port)
	return srv.ListenAndServe()
}
func main() {
	godotenv.Load()
	testHost := os.Getenv("TEST_HOST")
	fmt.Println(testHost)

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not found in the environment")
	}
	connection, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Can not connect to database", err)
	}

	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	cfg.DB = database.New(connection)

	app := &application{
		config: cfg,
		infoLog: infoLog,
		errorLog: errorLog,
		version: "1.0.0",
	}

	// starting the server 
	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)	
	}	
}