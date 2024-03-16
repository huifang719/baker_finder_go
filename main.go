package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type config struct {
	port int
	env string
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


	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config: cfg,
		infoLog: infoLog,
		errorLog: errorLog,
		version: "1.0.0",
	}

	// starting the server 
	err := app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)	
	}	
}