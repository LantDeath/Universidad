package main

import "fmt"

type Prenda struct {
	Tipo   string
	Talla  int
	Precio float64
}

// Factorial usando iteración
func factorialIterativo(n int) int {
	resultado := 1
	for i := 1; i <= n; i++ {
		resultado *= i
	}
	return resultado
}

// Factorial usando recursión
func factorialRecursivo(n int) int {
	if n == 0 {
		return 1 // Caso base: el factorial de 0 es 1
	}
	return n * factorialRecursivo(n-1) // Llamada recursiva
}

// Fibonacci usando iteración
func fibonacciIterativo(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Fibonacci usando recursión
func fibonacciRecursivo(n int) int {
	if n <= 1 {
		return n // Caso base
	}
	return fibonacciRecursivo(n-1) + fibonacciRecursivo(n-2) // Caso recursivo
}

func main() {
	numero := 4

	//Func factorial
	fmt.Printf("Factorial de %d (iterativo): %d\n", numero, factorialIterativo(numero))

	fmt.Printf("Factorial de %d (recursivo): %d\n", numero, factorialRecursivo(numero))

	//Func fibonacci
	fmt.Printf("Fibonacci de %d (iterativo): %d\n", numero, fibonacciIterativo(numero))

	fmt.Printf("Fibonacci de %d (recursivo): %d\n", numero, fibonacciRecursivo(numero))
}
