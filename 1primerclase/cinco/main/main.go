package main

import "fmt"

func main() {
	var resultado = miFuncion(2, 3, 2, 1, 2, 3, 4, 5, 6)
	fmt.Println(resultado)

}
func miFuncion(valores ...float64) float64 {

	var total float64
	for _, v := range valores {
		total += v

	}

	return total
}
