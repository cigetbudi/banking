package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	//ROUTE
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/getAll", getAllCustomers)

	//SERVER
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
