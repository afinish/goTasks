package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCraete(t *testing.T) {
	// Conta = make(map[string]Contact)
	CCounter = 0
	expected := Contact{"FName", "LName", "+11100", "e@mai.l", "Position"}
	actual := Create("FName", "LName", "+11100", "e@mai.l", "Position")
	assert.Equal(t, expected, actual, "Unexpected contact")
	assert.Equal(t, 1, CCounter, "Unexpected id counter value")
}
func TestUpdate(t *testing.T) {
	CCounter = 0
	Create("FName", "LName", "+11100", "e@mai.l", "Position")
	expected := Contact{"UptName", "LName", "+11100", "e@mai.l", "UptPosition"}
	actual := Update("1", "UptName", "LName", "+11100", "e@mai.l", "UptPosition")
	assert.Equal(t, expected, actual, "Unexpected updated contact values")
}
func TestGet(t *testing.T) {
	CCounter = 0
	Create("FName", "LName", "+11100", "e@mai.l", "Position")
	Create("2FName", "2LName", "2+11100", "2e@mai.l", "2Position")
	
	expected := Contact{"2FName", "2LName", "2+11100", "2e@mai.l", "2Position"}
	actual := Get("2")
	assert.Equal(t, expected, actual, "Unexpected contact")
	assert.Equal(t, 2, CCounter, "Unexpected id counter value")
}
func TestGetAll(t *testing.T) {
	CCounter = 0
	Create("1FName", "LName", "+11100", "e@mai.l", "Position")
	Create("2FName", "LName", "+11100", "e@mai.l", "Position")
	Create("3FName", "LName", "+11100", "e@mai.l", "Position")
	Create("4FName", "LName", "+11100", "e@mai.l", "Position")

	expected := map[string]Contact{
		"1": {
			"1FName", "LName", "+11100", "e@mai.l", "Position",
		},
		"2": {
			"2FName", "LName", "+11100", "e@mai.l", "Position",
		},
		"3": {
			"3FName", "LName", "+11100", "e@mai.l", "Position",
		},
		"4": {
			"4FName", "LName", "+11100", "e@mai.l", "Position",
		},
	}
	actual := GetAll()
	assert.Equal(t, expected, actual, "Unexpected contact list")
}
func TestDeleteContact(t *testing.T) {
	CCounter = 0
	Create("FName", "LName", "+11100", "e@mai.l", "Position")
	actual := DeleteContact("1")
	expected := Contact{}
	assert.Equal(t, expected, actual, "Contact might not be deletd")
}