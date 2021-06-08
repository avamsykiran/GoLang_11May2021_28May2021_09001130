package controllers

import (
	"net/http"
	"postgre-contacts-api/app/models"
	"postgre-contacts-api/app/services"

	"github.com/revel/revel"
)

type ContactController struct {
	*revel.Controller
}

func (c ContactController) GetContactsReqHandler() revel.Result {
	contacts, err := services.GetAllContacts()
	if err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.Render()
	}
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
	c, err := services.AddContact(contact)
	if err != nil {
		c.Response.Status = http.StatusInternalServerError
		return c.Render()
	}
	return c.RenderJSON(c)
}

func (c ContactController) UpdateContactByIdReqHandler(id int, contact models.Contact) revel.Result {

	if id == contact.ContactId {
		c, err := services.UpdateContact(contact)
		if err != nil {
			return c.NotFound("contact for the given id is not found")
		} else {
			return c.RenderJSON(c)
		}
	} else {
		c.Response.Status = http.StatusBadRequest
		return c.Render()
	}
}
