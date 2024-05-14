package main

import "fmt"

type Autor struct {
	Name     string
	Apellido string
}

func main() {
	a := Autor{
		Name:     "juan",
		Apellido: "Gutierrez",
	}

	l := Libro{
		Titulo:      "como en el lodo",
		Descripcion: "libro de suspenso",
		Autor:       a,
	}
	l.informacion()

}
func (a Autor) nombreCompleto() string {
	return fmt.Sprintf("%s, %s", a.Name, a.Apellido)
}

type Libro struct {
	Titulo      string
	Descripcion string
	Autor       Autor
}

func (l Libro) informacion() {
	fmt.Printf("Titulo: %s \nDescripcion: %s \nAutor: %s", l.Titulo, l.Descripcion, l.Autor.nombreCompleto())
}
