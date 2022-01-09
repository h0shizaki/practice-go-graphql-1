package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Config struct {
	Port        int
	Version     string
	Environment string
}

type Application struct {
	Config Config
	Logger *log.Logger
}

type Status struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func main() {

	var config Config
	flag.IntVar(&config.Port, "port", 8080, "Port of server")
	flag.StringVar(&config.Environment, "environment", "Develop", "Environment")
	flag.StringVar(&config.Version, "version", "1.0.0", "Version")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := Application{
		Config: config,
		Logger: logger,
	}

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := srv.ListenAndServe()
	fmt.Println("Sever is running on port ", config.Port)

	if err != nil {
		fmt.Println(err)
	}

}
