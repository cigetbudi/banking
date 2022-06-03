package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func main() {

	//ROUTE
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/getAll", getAllCustomers)

	//SERVER
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HALOOO")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Sigit", City: "Jakarta", ZipCode: "56161"},
		{Name: "Amel", City: "Kuningan", ZipCode: "66626"},
		{Name: "Tono", City: "Magelang", ZipCode: "451451"},
	}
	//addheader
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode((customers))
}
