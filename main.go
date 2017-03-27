package main

import (
	"cf-example-go-cassandra-app/models"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	DEFAULT_PORT = "8080"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	myRecord := models.GetRecord("hello")
	fmt.Fprintln(w, myRecord.Value)
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Printf("Warning, PORT not set. Defaulting to %+v\n", DEFAULT_PORT)
		port = DEFAULT_PORT
	}

	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Printf("ListenAndServe: ", err)
	}
}
