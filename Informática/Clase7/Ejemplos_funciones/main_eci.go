package main

import (
	"fmt"
	"math"
)

// 1. DEFINE UNA FUNCIÓN (devuelvePi) QUE DEVUELVA UN STRING Y UN VALOR DE TIPO float32 ASOCIADOS CON LA CONSTANTE Pi DEL PAQUETE math,
func devuelvePi() (string, float32) {
	return "Pi", math.Pi
}

// 2. DEFINE UNA FUNCIÓN (tomaDosArgsDevuelveDobleProducto) QUE TOME COMO ARGUMENTO DOS VALORES DE TIPO int Y QUE DEVUELVA EL DOBLE DE SU PRODUCTO
func tomaDosArgsDevuelveDobleProducto(x int, y int) int {

	return (x * y) * 2
}

// 3. DECLARA UNA FUNCIÓN QUE TOMARÁ UN ARGUMENTO ENTERO Y DEVOLVERÁ UN ENTERO
func la_funcion_devuelve_entero_8mil(a int) int {

	// 3.1 DECLARA DE FORMA ABREVIADA LA VARIABLE x CON UN VALOR DE TU GUSTO (EJ 45)
	b := 57

	// 3.2 HAZ QUE LA FUNCIÓN LA SUMA DE ESTA VARIABLE CON EL DEL ARGUMENTO
	return a + b
}

func main() {

	//Print de la primera función
	pi, numero_pi := devuelvePi()
	fmt.Println("El número correspondia a", pi, "es", numero_pi) // Llamamos el string y el float32 en lugar de llamar toda la función.

	// Print de la segunda función
	fmt.Println("El doble del producto de 53 y 17 es", tomaDosArgsDevuelveDobleProducto(53, 17))

	//Print de la tercera función
	fmt.Println("La suma de 57 y 33 es", la_funcion_devuelve_entero_8mil(33))
}
