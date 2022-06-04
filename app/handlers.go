package app

import (
	"banking/service"
	"encoding/json"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "Sigit", City: "Jakarta", ZipCode: "56161"},
	// 	{Name: "Amel", City: "Kuningan", ZipCode: "66626"},
	// 	{Name: "Tono", City: "Magelang", ZipCode: "451451"},
	// }
	//addheader

	customers, _ := ch.service.GetAllCustomer()
	//json
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode((customers))

	//XML
	// w.Header().Add("Content-Type", "application/xml")
	// xml.NewEncoder(w).Encode((customers))

}
