package services

import (
	"contacts-api/app/models"
)

var contactGroups []models.ContactGroup

func init() {
	//initialize some hypothetical data
	contactGroups = []models.ContactGroup{}
	contactGroups = append(contactGroups, models.ContactGroup{1, "Family"})
	contactGroups = append(contactGroups, models.ContactGroup{1, "Gym"})
	contactGroups = append(contactGroups, models.ContactGroup{1, "Work"})
}
