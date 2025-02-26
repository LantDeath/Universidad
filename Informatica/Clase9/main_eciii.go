package main

import (
	"fmt"
)

func pruebaBucleSencillo(total int) int {
	fmt.Println("Total pasado como argumento ", total)
	for i := 0; i < total; i++ {
		fmt.Println("Iteración ", i)
	}
	return total
}

func sumaN(total int) int {
	// Calcula la suma de los naturales entre 0 y total y devolverlo.
	var i int
	var suma int

	suma = 0

	for i = 0; i < total; i++ {
		suma = suma + i
	}
	return suma

}

func main() {

	var total int

	fmt.Println("Escribe el total de iteraciones ")
	fmt.Scanln(&total)
	_ = pruebaBucleSencillo(total)
	fmt.Println("La suma de los ", total, "primeros números naturales es", sumaN(total))

}
