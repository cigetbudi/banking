package app

import (
	"banking/domain"
	"banking/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	//WIRING
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	//ROUTE
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	//SERVER
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
