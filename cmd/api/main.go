package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type Config struct {
	Port        int
	Version     string
	Environment string
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world"))
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		currentStatus := Status{
			Status:      "Available",
			Environment: config.Environment,
			Version:     config.Version,
		}

		js, err := json.MarshalIndent(currentStatus, "", "\t")

		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)

	})

	fmt.Println("Server is running on port ", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)

}
