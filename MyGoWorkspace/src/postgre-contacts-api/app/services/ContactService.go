package services

import (
	"errors"
	"fmt"
	"postgre-contacts-api/app/models"
	"postgre-contacts-api/app/repos"
)

func GetAllContacts() (contacts []models.Contact, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
	}()
	return repos.ContactRepo.GetAllContacts(), nil
}

func GetContactById(cid int) (c models.Contact, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
	}()
	return repos.ContactRepo.GetContactById(cid), nil
}

func AddContact(contact models.Contact) (c models.Contact, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
	}()
	return repos.ContactRepo.AddContact(contact), nil
}

func UpdateContact(contact models.Contact) (c models.Contact, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
	}()
	return repos.ContactRepo.UpdateContact(contact), nil
}

func DeleteContact(cid int) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
	}()
	repos.ContactRepo.DelContact(cid)
	return nil
}
