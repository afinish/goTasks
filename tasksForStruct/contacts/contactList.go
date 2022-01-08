package contacts

import (
	"fmt"
	"strconv"
)

var cCounter int = 0
var contacts map[string]Contact

type Contact struct {
	FirstName 	string
	LastName 	string
	Phone 		string
	Email 		string
	Position 	string
}

func create(firstName, lastName, phone, email, position string) {
	cCounter++
	id := strconv.Itoa(cCounter)
	contacts[id] = Contact{
		firstName, 
		lastName,
		phone,
		email,
		position,
	}
}

func update(id, firstName, lastName, phone, email, position string) {
	contacts[id] = Contact{
		firstName, 
		lastName,
		phone,
		email,
		position,
	}
}

func get(id string) Contact {
	return contacts[id]
}

func getAll() map[string]Contact {
	return contacts
}

func deleteContact(id string) {
	delete(contacts, id)
}

func init() {
	contacts = make(map[string]Contact)
	
	create(
		"John",
		"Doe",
		"+998707777777",
		"johnny@example.com",
		"Team Lead",
	)

	create(
		"Jane",
		"Doe",
		"+998906669966",
		"janee@example.com",
		"Product Manager",
	)

	create(
		"Mark",
		"Marlin",
		"+781223748899",
		"markmarlin@example.com",
		"Client",
	)

	fmt.Println(get("2"))
	fmt.Println(getAll())

	update(
		"1",
		"John",
		"Doe",
		"+888707777777",
		"johnny@example.com",
		"Team Lead",
	)
	
	fmt.Println(get("1"))
	deleteContact("3")
	fmt.Println(getAll())
}
