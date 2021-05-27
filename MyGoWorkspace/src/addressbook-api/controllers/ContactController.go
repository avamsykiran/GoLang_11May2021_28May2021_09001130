package controllers

import (
	"addressbook-api/models"
	"addressbook-api/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetContactsReqHandler(w http.ResponseWriter, r *http.Request) {
	var contacts []models.Contact = services.GetAllContacts()
	json.NewEncoder(w).Encode(contacts)
}

func GetContactByIdReqHandler(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	id, _ := strconv.Atoi(pathVars["id"])

	contact, err := services.GetContactById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(contact)
	}
}

func DelContactByIdReqHandler(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	id, _ := strconv.Atoi(pathVars["id"])

	err := services.DeleteContact(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func AddContactReqHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var contact models.Contact
	err := json.Unmarshal(reqBody, &contact)
	if err == nil {
		services.AddContact(contact)
		json.NewEncoder(w).Encode(contact)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
	}
}

func UpdateContactByIdReqHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var contact models.Contact
	err := json.Unmarshal(reqBody, &contact)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	pathVars := mux.Vars(r)
	id, _ := strconv.Atoi(pathVars["id"])

	if id == contact.ContactId {
		err := services.UpdateContact(contact)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			json.NewEncoder(w).Encode(contact)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
