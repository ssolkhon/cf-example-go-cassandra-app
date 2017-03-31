package main

import (
	"cf-example-go-cassandra-app/cf"
	"cf-example-go-cassandra-app/db/cassandra"
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	DEFAULT_PORT   = "8080"
	DEFAULT_CONFIG = "./example_config.json"
)

func handleFavicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

func handleRequest(session *gocql.Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		myPath := strings.Split(r.URL.Path, "/")

		switch {
		case len(myPath) == 2:
			if myPath[1] != "" {
				err := cassandra.CreateTable(session, myPath[1])
				if err != nil {
					msg := "Error: " + err.Error()
					fmt.Fprintf(w, msg)
					log.Println(msg)
				} else {
					msg := "Successfully created collection " + myPath[1]
					fmt.Fprintf(w, msg)
					log.Printf(msg)
				}
			} else {
				msg := "Error: No parameters provided"
				fmt.Fprintf(w, msg)
			}
		case len(myPath) == 3:
			msg := "Coming soon..."
			fmt.Fprintf(w, msg)
		case len(myPath) == 4:
			err := cassandra.CreateRow(session, myPath[1], myPath[2], myPath[3])
			if err != nil {
				msg := "Error adding record: " + err.Error()
				log.Println(msg)
				fmt.Fprintf(w, msg)
			} else {
				msg := `Success, Added key/value: ` + myPath[2] + `/` + myPath[3]
				log.Println(msg)
				fmt.Fprintf(w, msg)
			}
		default:
			msg := "Error: Too many parameters provided"
			fmt.Fprintf(w, msg)
		}
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
			log.Printf("Error loading default config file: %v\n", err)
			os.Exit(1)
		}
		servicesRaw = file
	}
	// Set myServices
	err := json.Unmarshal(servicesRaw, &myServices)
	if err != nil {
		log.Printf("json.Unmarshal() error: %v\n", err)
		os.Exit(1)
	}
	// Set Cassandra connection
	mySession, err := cassandra.GetSession(myServices)
	if err != nil {
		log.Printf("Error loading default config file: %v\n", err)
		os.Exit(1)
	}
	defer mySession.Close()

	// Handle Requests
	http.HandleFunc("/favicon.ico", handleFavicon)
	http.HandleFunc("/", handleRequest(mySession))
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Printf("ListenAndServe: ", err)
	}
}
