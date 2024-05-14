package main

import "fmt"

const (
	suma   = "+"
	resta  = "-"
	multip = "*"
	divis  = "/"
)

func main() {

	var resultado = miFuncion(suma, 3, 2, 1, 2, 3, 4, 5, 6)
	fmt.Println(resultado)
	var resultado2 = miFuncion(resta, 3, 2, 1, 2, 3, 4, 5, 6)
	fmt.Println(resultado2)
	var resultado3 = miFuncion(multip, 3, 2, 1, 2, 3, 4, 5, 6)
	fmt.Println(resultado3)
	var resultado4 = miFuncion(divis, 3, 2, 1, 2, 3, 4, 5, 6)
	fmt.Println(resultado4)

}

func miFuncion(valor string, valores ...float64) float64 {
	var total float64
	if valor == "+" {
		for _, v := range valores {
			total += v

		}
	} else if valor == "-" {
		for _, v := range valores {
			total -= v

		}
	} else if valor == "*" {
		for _, v := range valores {
			total = valores[0]
			total *= v

		}
	} else if valor == "/" {
		for _, v := range valores {
			total = valores[0]
			total /= v

		}
	}

	return total
}
