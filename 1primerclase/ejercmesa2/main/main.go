package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
var empleados int = 0

func main() {
	fmt.Println("Benjamin tiene: ", employees["Benjamin"])
	for _, v := range employees {
		if v > 21 {
			empleados++
		}

	}
	fmt.Println("los mayores a 21 son ", empleados)
	employees["Federico"] = 25
	fmt.Println("Se agrego un nuevo empleado", employees["Federico"])
	fmt.Println("nueva lista", employees)

	delete(employees, "Pedro")
	fmt.Println("se elimino un empleado: ")
	fmt.Println("nueva lista ", employees)
}
