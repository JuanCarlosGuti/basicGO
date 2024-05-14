package main

import (
	"errors"
	"fmt"
)

func main() {
	opName := "average"
	opVar, err := operation(opName)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("%s is %v", opName, opVar(2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func operation(operacion string) (func(...int) int, error) {
	const (
		minimum = "minimum"
		maximum = "maximum"
		average = "average"
	)
	switch operacion {
	case minimum:
		return returnMinimo, nil
	case maximum:
		return returnMaximo, nil
	case average:
		return returnAvarage, nil
	default:
		return nil, errors.New("no existe esa operación")
	}

}

func returnMinimo(valor ...int) int {
	var nuevo = valor[0]
	for _, v := range valor {
		if v < nuevo {
			nuevo = v
		}
	}
	return nuevo

}

func returnMaximo(valor ...int) int {
	var nuevo = valor[0]
	for _, v := range valor {
		if v > nuevo {
			nuevo = v
		}
	}
	return nuevo
}
func returnAvarage(valor ...int) int {
	largo := len(valor)
	media := 0

	for _, v := range valor {
		media += v
	}

	return media / largo

}
