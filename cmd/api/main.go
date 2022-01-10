package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/models"
	"time"

	_ "github.com/lib/pq"
)

type Config struct {
	Port        int
	Version     string
	Environment string
	DB          struct {
		DSN string
	}
}

type Application struct {
	Config Config
	Logger *log.Logger
	Models models.Models
}

type Status struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func main() {

	var config Config
	flag.IntVar(&config.Port, "port", 8080, "Port of server")
	flag.StringVar(&config.Environment, "environment", "Development", "Environment")
	flag.StringVar(&config.Version, "version", "1.0.0", "Version")
	flag.StringVar(&config.DB.DSN, "dsn", "postgres://postgres:22194@localhost/mydb?sslmode=disable", "PostgresQL")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(config)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	app := Application{
		Config: config,
		Logger: logger,
		Models: models.NewModel(db),
	}

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = srv.ListenAndServe()
	fmt.Println("Sever is running on port ", config.Port)

	if err != nil {
		fmt.Println(err)
	}

}

func openDB(config Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DB.DSN)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
