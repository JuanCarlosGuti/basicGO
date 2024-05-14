package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	DNI   int
	Fecha string
}

func main() {
	s := Student{
		Name:  "John Doe",
		Age:   42,
		DNI:   4256478995,
		Fecha: "06/04/2024",
	}
	s.Detalle()

}

func (s *Student) Detalle() {
	fmt.Printf("%s is %d years old his/her DNI is %d and inscripcion date is %s .\n", s.Name, s.Age, s.DNI, s.Fecha)
}
