package controllers

import (
	"contacts-api/app/models"
	"contacts-api/app/services"

	"github.com/revel/revel"
)

type ContactController struct {
	*revel.Controller
}

func (c ContactController) GetContactsReqHandler() revel.Result {
	var contacts []models.Contact = services.GetAllContacts()
	c.RenderJson(contacts)
}

func (c ContactController) GetContactByIdReqHandler(id int) revel.Result {

	contact, err := services.GetContactById(id)
	if err != nil {
		c.NotFound("contact for the given id is not found")
	} else {
		c.RenderJson(contact)
	}
}

func (c ContactController) DelContactByIdReqHandler(id int) revel.Result {
	err := services.DeleteContact(id)
	if err != nil {
		c.NotFound("contact for the given id is not found")
	} else {
		c.Ok("contact deleted")
	}
}

func (c ContactController) AddContactReqHandler(contact models.Contact) revel.Result {
	services.AddContact(contact)
	c.RenderJson(contact)
}

func (c ContactController) UpdateContactByIdReqHandler(id int, contact models.Contact) revel.Result {

	if id == contact.ContactId {
		err := services.UpdateContact(contact)
		if err != nil {
			c.NotFound("contact for the given id is not found")
		} else {
			c.RenderJson(contact)
		}
	} else {
		c.BadRequest("contact passed and contact intendd to update do not match")
	}
}
