package main

import "fmt"

type Gato struct {
	Name string
}

func main() {
	fmt.Println("la funcion se esta iniciando")
	ejemplo()
	fmt.Println("la funcion se esta terminando")
} /*func ejemplo() {
	panic("oooo")

}*/

func ejemplo() {
	s := &Gato{"Misho"}
	s = nil
	s.Hablar()
}
func (s *Gato) Hablar() {
	fmt.Println(s.Name, "el gato hace miau")
}
