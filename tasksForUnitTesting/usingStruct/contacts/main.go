package main

import (
	"fmt"
	"strconv"
)

var CCounter int = 0
var Contacts map[string]Contact = make(map[string]Contact)
// Contacts 

type Contact struct {
	FirstName 	string
	LastName 	string
	Phone 		string
	Email 		string
	Position 	string
}

func Create(firstName, lastName, phone, email, position string) Contact {
	CCounter++
	id := strconv.Itoa(CCounter)
	Contacts[id] = Contact{
		firstName, 
		lastName,
		phone,
		email,
		position,
	}
	return Contacts[id]
}

func Update(id, firstName, lastName, phone, email, position string) Contact {
	Contacts[id] = Contact{
		firstName, 
		lastName,
		phone,
		email,
		position,
	}
	return Contacts[id]
}

func Get(id string) Contact {
	return Contacts[id]
}

func GetAll() map[string]Contact {
	return Contacts
}

func DeleteContact(id string) Contact {
	delete(Contacts, id)
	return Contacts[id]
}

func main() {
	
	
	fmt.Println(Create(
		"John",
		"Doe",
		"+998707777777",
		"johnny@example.com",
		"Team Lead",
	))

	fmt.Println(Create(
		"Jane",
		"Doe",
		"+998906669966",
		"janee@example.com",
		"Product Manager",
	))

	fmt.Println(Create(
		"Mark",
		"Marlin",
		"+781223748899",
		"markmarlin@example.com",
		"Client",
	))

	fmt.Println(Get("2"))
	fmt.Println(GetAll())

	fmt.Println(Update(
		"1",
		"John",
		"Doe",
		"+888707777777",
		"johnny@example.com",
		"Team Lead",
	))
	
	fmt.Println(Get("1"))
	fmt.Println(DeleteContact("3"))
	fmt.Println(GetAll())
}