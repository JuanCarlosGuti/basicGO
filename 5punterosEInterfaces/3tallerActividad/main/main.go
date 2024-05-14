package main

import "fmt"

func main() {
	usuario1 := Usuarios{
		Nombre:     "Juan",
		Apellido:   "Gutierrez",
		Edad:       37,
		Correo:     "juan@gmail.com",
		Contrasena: "860406",
	}
	fmt.Println(usuario1)
	usuario1.cambiarApellido("huerfano")
	fmt.Println("Cambio Apellido", usuario1)

}

type Usuarios struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

func (u *Usuarios) cambiarNombre(nombre string) {
	u.Nombre = nombre

}

func (u *Usuarios) cambiarApellido(apellido string) {
	u.Apellido = apellido
}
func (u *Usuarios) cambiarEdad(edad int) {
	u.Edad = edad
}
func (u *Usuarios) cambiarCorreo(correo string) {
	u.Correo = correo
}

func (u *Usuarios) cambiarContraseña(contrasenia string) {
	u.Contrasena = contrasenia

}
