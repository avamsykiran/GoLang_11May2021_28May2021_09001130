package controllers

import (
	"contacts-api/app/models"
	"contacts-api/app/services"
	"net/http"

	"github.com/revel/revel"
)

type ContactController struct {
	*revel.Controller
}

func (c ContactController) GetContactsReqHandler() revel.Result {
	var contacts []models.Contact = services.GetAllContacts()
	return c.RenderJSON(contacts)
}

func (c ContactController) GetContactByIdReqHandler(id int) revel.Result {

	contact, err := services.GetContactById(id)
	if err != nil {
		return c.NotFound("contact for the given id is not found")
	} else {
		return c.RenderJSON(contact)
	}
}

func (c ContactController) DelContactByIdReqHandler(id int) revel.Result {
	err := services.DeleteContact(id)
	if err != nil {
		return c.NotFound("contact for the given id is not found")
	} else {
		c.Response.Status = http.StatusOK
		return c.Render()
	}
}

func (c ContactController) AddContactReqHandler(contact models.Contact) revel.Result {
	services.AddContact(contact)
	return c.RenderJSON(contact)
}

func (c ContactController) UpdateContactByIdReqHandler(id int, contact models.Contact) revel.Result {

	if id == contact.ContactId {
		err := services.UpdateContact(contact)
		if err != nil {
			return c.NotFound("contact for the given id is not found")
		} else {
			return c.RenderJSON(contact)
		}
	} else {
		c.Response.Status = http.StatusBadRequest
		return c.Render()
	}
}
