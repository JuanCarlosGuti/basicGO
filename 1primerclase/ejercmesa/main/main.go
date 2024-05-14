package main

import "fmt"

var alumnos = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
var nuevostudiante = "Gabriela"

func main() {
	if nuevostudiante == "" {
		fmt.Println("Nuevo studiante no existe")
		fmt.Println("lista estudiantes: ", alumnos)
	} else if nuevostudiante != "" {
		alumnos = append(alumnos, nuevostudiante)
		fmt.Println("hay un nuevo estudiante")
		fmt.Println("nueva lista de estudiantes: ", alumnos)
	}

}
