package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       int
	Description string
	Category    string
}

func main() {
	p1 := Product{ID: 1, Name: "Producto 1", Price: 10, Description: "Descripción del Producto 1", Category: "Categoría 1"}
	p2 := Product{ID: 2, Name: "Producto 2", Price: 20, Description: "Descripción del Producto 2", Category: "Categoría 2"}
	p3 := Product{ID: 3, Name: "Producto 3", Price: 30, Description: "Descripción del Producto 3", Category: "Categoría 3"}

	p1.save()
	p2.save()
	p3.save()

	Product{}.GetAll()

}

var Products []Product

func (p Product) save() {
	Products = append(Products, p)

}

func (p Product) GetAll() {
	for _, pr := range Products {
		fmt.Printf("Product ID: %d\nName: %s\nPrice: %d\nDescription: %s\nCategory: %s\n", pr.ID, pr.Name, pr.Price, pr.Description, pr.Category)
	}
}
