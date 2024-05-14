package main

import (
	"fmt"
)

func main() {
	var v int = 50
	var p *int
	p = &v

	fmt.Println("la direccion de memoria es ", &v)
	fmt.Println("la direccion de memoria es ", p)
	fmt.Println("el valor de la variable es", *p)
}
