package generator

import (
	"multi-insert/models"

	"github.com/jackc/fake"
)

func GenerateDummyContacts(numberOfContacts int) models.Contacts {
	var contacts models.Contacts
	for i := 0; i < numberOfContacts; i++ {
		contacts = append(contacts, models.Contact{
			FirstName: fake.FirstName(),
			LastName:  fake.LastName(),
			Email:     fake.EmailAddress(),
		})
	}

	return contacts
}
