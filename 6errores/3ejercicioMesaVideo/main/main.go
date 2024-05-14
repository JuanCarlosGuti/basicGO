package main

import (
	"fmt"
	"strconv"
)

func main() {
	var mensaje string
	fmt.Scan(&mensaje)
	fmt.Println("este fue el mensaje: ", mensaje)

	var (
		nombre   string
		apellido string
		telefno  int
	)
	fmt.Scanf("%s %s %d", &nombre, &apellido, &telefno)

	fmt.Println(nombre, apellido, telefno)

	intVar := 342
	intStrvar2 := strconv.Itoa(intVar)
	fmt.Println(intStrvar2)

}
