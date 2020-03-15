package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/bygui86/go-cassandra/cassandra"
	"github.com/bygui86/go-cassandra/users"
	"github.com/gorilla/mux"
)

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func heartBeatHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}

func main() {
	// Cassandra connection
	CassandraSession := cassandra.Session
	defer CassandraSession.Close()

	// REST APIs handling
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", heartBeatHandler)
	router.HandleFunc("/users/new", users.PostHandler)
	router.HandleFunc("/users", users.GetAllHandler)
	router.HandleFunc("/users/{user_uuid}", users.GetByUuidHandler)
	var port int = 8080
	fmt.Println("Server listen on port", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
