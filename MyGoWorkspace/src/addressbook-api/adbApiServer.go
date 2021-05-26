package main

import (
	"addressbook-api/models"
	"addressbook-api/services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getContactsReqHandler(w http.ResponseWriter, r *http.Request) {
	var contacts []models.Contact = services.GetAllContacts()
	json.NewEncoder(w).Encode(contacts)
}

func main() {
	fmt.Println("REST API - Contact")

	adbRouter := mux.NewRouter().StrictSlash(true)

	adbRouter.HandleFunc("/contacts", getContactsReqHandler)

	log.Fatal(http.ListenAndServe(":9999", adbRouter))
}
