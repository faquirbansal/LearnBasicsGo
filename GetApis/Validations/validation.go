package main

import "errors"

type PersonDataStruct struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	PhoneNo    int    `json:"phoneNo"`
	Gendar     string `json:"gendar"`
	Disability bool   `json:"disability"`
}

// Validate the incoming item payload
func ValidatePerson(person PersonDataStruct) error {
	if person.ID == 0 {
		return errors.New("id is required and should be greater than 0")
	}
	if person.Name == "" {
		return errors.New("name is required and cannot be empty")
	}
	if person.PhoneNo == 10 {
		return errors.New("phone no is required and should have 10digits")
	}
	return nil
}
