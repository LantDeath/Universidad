package main

import (
	"fmt"
)

// Dado que tenemos tiempo y hemos completado los otros puntos intentaré solucionar este punto aunque se acote la nota final a 10.

// Creamos la función Factorial
func Factorial(n int) int {
	resultado := 1
	for i := 1; i <= n; i++ {
		resultado *= i
	}
	return resultado
}

func main() {
	// Creamos una variable donde pondremos el número que querramos saber su factorial.
	var num int
	fmt.Println("Escriba un número para obtener su factorial:")
	fmt.Scanln(&num)
	fmt.Println("El factorial de ", num, "es", Factorial(num))

}
