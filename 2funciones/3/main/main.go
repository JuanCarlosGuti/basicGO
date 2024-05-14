package main

import (
	"fmt"
)

func main() {
	opName := "clasea"
	opVar, err := liquidar(opName)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf(" el empleado que es de %s se le liquido %v", opName, opVar(6000))
}

func liquidar(operacion string) (func(valor float32) float32, error) {
	const (
		clasea = "clasea"
		claseb = "claseb"
		clasec = "clasec"
	)
	switch operacion {
	case clasea:
		return a, nil
	case claseb:
		return b, nil
	case clasec:
		return c, nil
	}
	return nil, fmt.Errorf("Invalid operacion")

}

func a(valor float32) float32 {
	minuto := float32(3000 / 60)
	salarioMensual := minuto * valor
	salarioBonificacion := salarioMensual * 1.5

	return salarioBonificacion
}

func b(valor float32) float32 {
	minuto := float32(1500 / 60)
	salarioMensual := minuto * valor
	salarioBonificacion := salarioMensual * 1.2

	return salarioBonificacion
}
func c(valor float32) float32 {
	minuto := float32(1000 / 60)
	salarioMensual := minuto * valor

	return salarioMensual
}
