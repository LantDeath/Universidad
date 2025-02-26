package main

import (
	"fmt"
)

/*
func leePEntero(pi *int) int {
	fmt.Scanln(pi)
	return *pi
}

func leeEntero() int {
	var aux_return int

	fmt.Scanln(&aux_return)
	return aux_return
}

func pruebaLectura() {
	var ent1, ent2, ent3, ent4 int
	var f321 float32
	var s1 string
	var b1 bool

	// LECTURA BÁSICA DE CUALQUIER TIPO DE DATOS. CUIDADO
	// RECUERDA QUE DEBES EJECUTAR FUERA DEL ENTORNO
	// go mod main_ecii
	// go mod tidy
	// go run main_ecii.go
	fmt.Println("Se le va a solicitar que escribe 3 datos ")

	fmt.Println("Escribe un número real (xxx.xxxx) ")
	fmt.Scanln(&f321)
	fmt.Println("Escribe un string ")
	fmt.Scanln(&s1)
	fmt.Println("Escribe un valor booleano ")
	fmt.Scanln(&b1)
	fmt.Println("Resultados leidos ", f321, s1, b1)

	// Ejemplo sobre cómo leer enteros directamente y dentro de funciones

	fmt.Println("Se le va a solicitar que ecribe 3 enteros ")

	fmt.Println("Escribe el primer entero ")
	fmt.Scanln(&ent1)

	fmt.Println("Escribe el segundo entero")
	ent2 = leeEntero()

	// ESTAS SON OPCIONES AVANZADAS PARA AQUELLOS QUE QUIERAN PENSARLO UN POCO MÁS
	fmt.Println("Escribe el tercer entero ")
	ent4 = leePEntero(&ent3)

	fmt.Println("Este es el resultado de lo escrito ", ent1, ent2, ent3, ent4)
}

func pruebaCondicionalSimple() {
	var num1 int

	// EJEMPLO DE CONDICIONAL SIMPLE
	fmt.Println("Escribe un numero entero")
	fmt.Scanln(&num1)
	if num1 < 0 {
		fmt.Println(num1, "es negativo")
	} else {
		fmt.Println(num1, "no es negativo")
	}

}

func pruebaCondicionalAnidada() {
	var num1 int

	// EJEMPLO DE CONDICIONAL SIMPLE
	fmt.Println("Escribe un numero entero")
	fmt.Scanln(&num1)
	if num1 < 0 {
		fmt.Println(num1, "es negativo")
		if num1%2 == 0 {
			fmt.Println(num1, "y par")
		}
	} else {
		fmt.Println(num1, "no es negativo")
		if num1%2 != 0 {
			fmt.Println(num1, "e impar")
		}
	}

}


// ESTA ES PARA LOS MÁS AVANZADOS, RECUERDA SIEMPRE PUEDES SUSTITUIR UN switch POR if-then-elses ANIDADOS
func pruebaCondicionalMultiple() {
	var num int

	fmt.Println("Escribe un numero entero")
	fmt.Scanln(&num)

	switch {
	case num < 5:
		fmt.Println("Caso < 5", num)
		// PROBAR EL EFECTO fallthrough
		// fallthrough
	case num > 2:
		fmt.Println("Caso > 2", num)
		// PROBAR EL EFECTO fallthrough
		// fallthrough
	default:
		fmt.Println("defecto")
	}

}

*/

func ejCondicionalAnidada() {
	// 1. LEE DOS VARIABLES DE TIPO ENTERO
	var num1 int
	var num2 int

	fmt.Println("Escribe un número entero:")
	fmt.Scanln(&num1)
	fmt.Println("Escribe otro número entero:")
	fmt.Scanln(&num2)
	// 2. DEFINE VARIABLES bool PARA LAS CONDICIONES n1 mayor que n2, n1 par, n2 par

	// Condicional num1 mayor que num2.
	if num1 > num2 {
		fmt.Println(num1, "es mayor que", num2)
	} else {
		fmt.Println(num1, "es menor que", num2)
	}

	// Condicional n1 par.
	if num1%2 == 0 {
		fmt.Println(num1, "es un número par.")
	} else {
		fmt.Println(num1, "no es un número par.")
	}

	// Condicional n2 par.
	if num2%2 == 0 {
		fmt.Println(num2, "es un número par.")
	} else {
		fmt.Println(num2, "no es un número par.")
	}

	if num1 > num2 {
		fmt.Println(num1, "es menor que", num2)
		if num1%2 == 0 {
			fmt.Println(num1, "es par, y el producto de", num1, "y", num2, "es", num1*num2)
		} else {
			fmt.Println(num1, "no es un número par")
		}
	} else if num2%2 == 0 {
		fmt.Println(num2, "es par, y la suma entre", num1, "y", num2, "es", num1+num2)
	} else {
		fmt.Println("El resultado es 0")
	}

	switch {
	case num1 > num2:
		{
			if num1%2 == 0 {
				fmt.Println("El resultado es:", num1*num2)
			}
			if num1%2 != 0 {
				fmt.Println("El resultado es 0.")
			}
		}
	case num1 < num2:
		{
			if num2%2 == 0 {
				fmt.Println("La suma de", num1, "y", num2, "es", num1+num2)
			}
		}
	default:
		{
			fmt.Println("El resultado es 0.")
		}
	}

}

// 3. UTILIZA UNA ESTRUCTURA CONDICIONAL PARA EL SIGUIENTE ALGORITMO
// SI n1 > n2 , en este caso si n1 es par el resultado es el producto de n1 y n2
// En otro caso si n2 es par el resultado es la suma
// En cualquier otro caso el resultado es 0

func main() {

	//pruebaLectura()
	//pruebaCondicionalSimple()
	//pruebaCondicionalAnidada()
	//pruebaCondicionalMultiple()

	// EJERCICIO

	ejCondicionalAnidada()

}
