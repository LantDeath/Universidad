package main

import (
	"fmt"
)

// Creamos la función con sus condicionales.
func bisiesto(ano int) bool {
	if ano%400 == 0 {
		return true
	} else if ano%4 == 0 && ano%100 != 0 {
		return true
	}
	return false
}

func main() {
	// creamos la variable num1.
	var num1 int

	// Hacemos el llamado.
	fmt.Println("Introduza el año:")
	fmt.Scanln(&num1)
	fmt.Println("¿Es bisiesto", num1, "?", bisiesto(num1))
}
