package main

import (
	"cf-example-go-cassandra-app/cf"
	"cf-example-go-cassandra-app/record"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	DEFAULT_PORT   = "8080"
	DEFAULT_CONFIG = "./example_config.json"
)

func DefaultHandler(s *cf.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		myRecord := record.GetRecord("hello")

		fmt.Fprintf(w, "Record: %+v\n", myRecord.Value)
		fmt.Fprintf(w, "Name: %+v\n", s.Cassandra[0].Name)
	}
}

func main() {
	var port string
	servicesRaw := []byte(os.Getenv("VCAP_SERVICES"))
	myServices := &cf.Services{}

	// Check port
	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Printf("Warning, PORT not set. Defaulting to %+v\n", DEFAULT_PORT)
		port = DEFAULT_PORT
	}
	// Check services
	if len(servicesRaw) == 0 {
		log.Printf("Warning, VCAP_SERVICES not set. Defaulting to %+v\n", DEFAULT_CONFIG)
		file, err := ioutil.ReadFile(DEFAULT_CONFIG)
		if err != nil {
			fmt.Printf("Error loading default config file: %v\n", err)
			os.Exit(1)
		}
		servicesRaw = file
	}
	// Set myServices
	err := json.Unmarshal(servicesRaw, &myServices)
	if err != nil {
		fmt.Printf("json.Unmarshal() error: %v\n", err)
		os.Exit(1)
	}
	// Handle Requests
	http.HandleFunc("/", DefaultHandler(myServices))
	http.ListenAndServe(":"+port, nil)

}
