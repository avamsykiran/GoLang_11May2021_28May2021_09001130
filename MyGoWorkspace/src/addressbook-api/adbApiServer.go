package main

import (
	"addressbook-api/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("REST API - Contact")

	adbRouter := mux.NewRouter().StrictSlash(true)

	adbRouter.HandleFunc("/contacts", controllers.GetContactsReqHandler).Methods("GET")
	adbRouter.HandleFunc("/contacts", controllers.AddContactReqHandler).Methods("POST")
	adbRouter.HandleFunc("/contacts/{id}", controllers.GetContactByIdReqHandler).Methods("GET")
	adbRouter.HandleFunc("/contacts/{id}", controllers.UpdateContactByIdReqHandler).Methods("PUT")
	adbRouter.HandleFunc("/contacts/{id}", controllers.DelContactByIdReqHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9999", adbRouter))
}
