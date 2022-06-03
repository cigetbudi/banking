package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//ROUTE
	http.HandleFunc("/greet", greet)

	//SERVER
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HALOOO")
}
