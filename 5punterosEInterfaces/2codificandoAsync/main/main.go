package main

import "fmt"

type Compañia struct {
	Nombre string
	Puesto string
}
type Empleado struct {
	Nombre   string
	Apellido string
	Compañia Compañia
}

func (e Empleado) informacion() {
	fmt.Printf("\nEmpleado: %s %s \nPuesto: %s\nCompañia: %s", e.Nombre, e.Apellido, e.Compañia.Puesto, e.Compañia.Nombre)

}
func (c *Compañia) CambiarPuesto(nuevoPuesto string) {
	c.Puesto = nuevoPuesto

}

func main() {
	empleado := Empleado{
		Nombre:   "Juan",
		Apellido: "Gutierrez",
		Compañia: Compañia{
			Nombre: "Lifting",
			Puesto: "Dev jr",
		},
	}
	empleado.informacion()
	empleado.Compañia.CambiarPuesto("Product Owner")
	empleado.informacion()
}
