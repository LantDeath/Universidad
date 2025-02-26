package main

import (
	"fmt"
	"math"
	// MATEMATICAS "math"
)

func conversionesExplicitas() {
	var i16 int16
	var i64 int64
	var f64 float64
	var pf64 *float64

	i16 = 4
	i64 = int64(i16 * 2)

	fmt.Println("Valor de i16 ", i16, " y de i64 ", i64, " y de f64 ", f64, " y de pf64 ", pf64)
}

func expresionesAritmeticas() {
	var i64 int64
	var f64 float64

	i64 = 13

	f64 = float64(i64 % 2)

	fmt.Println("Valor de i64 ", i64, " y de i64 % 2 ", f64)

	i64 *= 2

	f64 = float64(i64 % 2)

	fmt.Println("Valor de i64 ", i64, " y de i64 % 2 ", f64)

}

func expresionesRelacion() {
	var i64 int64
	var f64 float64
	var b1, b2, b3 bool

	i64 = 13
	f64 = float64(i64 % 2)
	b1 = (f64 == 0)
	b2 = (f64 != 0)
	b3 = b1 || b2 //OR lógico o disyunción
	fmt.Println("Valor de i64 ", i64, " y de i64 % 2 ", f64, " de b1 ", b1, " de b2 ", b2, " y de b3 ", b3)

}

func ejemplosMath() {
	var i64 int64
	var f64 float64

	i64 = 13
	f64 = float64(i64 % 2)

	fmt.Println("Valor de i64 ", i64, " y de i64 % 2 ", f64)
	fmt.Println("Valor de pi ", math.Pi, " y de phi ", math.Phi)

}

func ejerciciosExpresiones() {
	// 1. DEFINE UNA VARIABLE REAL DE TAMAÑO 64 (LLÁMALA superficie) Y OTRA DE TAMAÑO 32 (LLÁMALA radio)
	var superficie float64
	var radio float32

	// 2. ASIGNALE A LA VARIABLE radio EL VALOR 12
	radio = 12

	// 3. ASÍGNALE A LA VARIABLE superficie EL VALOR DE LA SUPERFICIE DEL CÍDULO DE RADIO radio
	superficie = math.Pi * math.Pow(float64(radio), 2)

	// 4. MUESTRA UN MENSAJE EN EL QUE INFORMES DEL RESULTADO
	fmt.Println("El valor de la superficies es:", superficie)

	// 5. DECLARA UNA VARIABLE PUNTERO A LA VARIABLE radio. LLÁMALA p_radio
	var p_radio *float32
	p_radio = &radio

	// 6. ESCRIBE UNA SENTENCIA MEDIANTE LA QUE MODIFIQUES EL VALOR DE radio AL DOBLE USANDO SÓLO LA VARIABLE PUNTERO
	*p_radio *= 2

	// 7. COPIA LAS SENTENCIAS MEDIANTE LAS QUE CALCULASTE LA SUPERFICIE Y MOSTRABAS SU VALOR
	superficie = math.Pi * math.Pow(float64(radio), 2)
	fmt.Println("El valor de la ahora superficies es:", superficie)

}

func ejemplosComplex() {

	var c641 complex64 = complex(12, 2.2)

	fmt.Println("Valor del complejo c641 + 3+3i", c641+(3+3i))
}

func ejemplosPunteros() {

	var i1 int
	var pi1 *int
	var b1 bool
	var pb1 *bool

	i1 = 100
	fmt.Println("Valor de pi1 ", i1)
	pi1 = &i1

	b1 = true
	pb1 = &b1

	fmt.Println("Valor de i1 ", i1)
	*pi1 += 1
	fmt.Println("Se modifica i1 desde pi1, su valor es i1 ", i1)

	fmt.Println("Valor de b1 ", b1)
	*pb1 = (*pb1) && false
	fmt.Println("Se modifica b1 desde pb1, su valor es b1 ", b1)
}

func main() {

	conversionesExplicitas()
	expresionesAritmeticas()
	expresionesRelacion()
	ejemplosMath()
	ejemplosComplex()
	ejemplosPunteros()

	ejerciciosExpresiones()

}
