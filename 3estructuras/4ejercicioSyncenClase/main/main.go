package main

import (
	"fmt"
)

type Person struct {
	ID          string
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       string
	Position string
	Person   Person
}

func main() {

	personOne := Person{
		ID:          "1",
		Name:        "John Doe",
		DateOfBirth: "1999-01-01",
	}

	employeeOne := Employee{
		ID:       "1",
		Position: "Developer",
		Person:   personOne,
	}
	employeeOne.PrintEmployee()
}

func PrintEmployee() {

}

func (e Employee) PrintEmployee() {
	fmt.Printf("El empleado con ID %s posiión %s y nombre %s con fecha de nacimiento %s .\n", e.ID, e.Position, e.Person.Name, e.Person.DateOfBirth)
}
