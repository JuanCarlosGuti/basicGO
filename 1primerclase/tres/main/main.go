package main

import "fmt"

func main() {
	meses := []string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio"}
	Obtener(meses)

}

func Obtener(meses []string) {
	for _, mes := range meses {
		if mes == "enero" || mes == "febrero" || mes == "marzo" {
			fmt.Println("verano")
		} else {
			fmt.Println("no")
		}
	}
}
