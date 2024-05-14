package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	p := Person{
		Name: "John",
		Age:  20,
	}
	p.Description()

}

func (p *Person) Description() {
	fmt.Printf("%s is %d years old.\n", p.Name, p.Age)
}
