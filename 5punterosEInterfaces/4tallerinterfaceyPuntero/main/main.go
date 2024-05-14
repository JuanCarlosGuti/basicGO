package main

import "fmt"

func main() {
	producto1 := nuevoProducto("pequenio", 25000)
	fmt.Println("Producto1 is", producto1.Precio())

	producto2 := nuevoProducto("mediano", 45000)
	fmt.Println("Producto2 is", producto2.Precio())

	producto3 := nuevoProducto("grande", 75000)
	fmt.Println("Producto3 is", producto3.Precio())

	producto4 := nuevoProducto("pequenios", 585000)
	fmt.Println("Producto4 is", producto4.Precio())

}

type Producto interface {
	Precio() float64
}

type pequenio struct {
	precio float64
}
type mediano struct {
	precio float64
}
type grande struct {
	precio float64
}

func (p pequenio) Precio() float64 {
	return p.precio
}
func (m mediano) Precio() float64 {
	return m.precio * 1.03
}
func (g grande) Precio() float64 {
	return g.precio*1.06 + 2500
}

func nuevoProducto(tipo string, precio float64) Producto {
	switch tipo {
	case "pequenio":
		return pequenio{precio: precio}
	case "mediano":
		return mediano{precio: precio}
	case "grande":
		return grande{precio: precio}
	default:
		panic("Invalid tipo")
	}

}
