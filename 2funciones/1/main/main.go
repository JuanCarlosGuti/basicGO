package main

import (
	"fmt"
)

func main() {
	total := pagoImpuesto(50000)
	fmt.Println("El impuesto es: ", total)

	total2 := promedio(25, 35, 45, 88)
	fmt.Println("El promedio es: ", total2)

	//	resultado := suma(4, 8)
	//	fmt.Println("El resultado: ", resultado)

}

/*
	 func suma(a, b int) int {
		return a + b
	}
*/
func pagoImpuesto(b float32) float32 {
	var resultado float32
	if b >= 150000 {
		resultado = b - (b * 0.27)
		return resultado
	} else if b >= 50000 {
		resultado = b - (b * 0.17)
		return resultado
	} else {
		resultado = b
		return resultado
	}
}

func promedio(valor ...float32) float32 {
	var media float32
	var resultado float32
	largo := len(valor)
	for _, v := range valor {
		resultado += v

	}
	media = resultado / float32(largo)
	return media

}
