package main

import (
	"fmt"
)

func pruebaArraysSlices() {
	// ARRAY Y SLICES

	var array_float = [...]float64{99: 100.0}
	var array_float2 = [...]float64{199: 200, 99: 100.0}

	fmt.Println("Valor en la posicion del array de floats 99:",
		array_float[98])
	fmt.Println("Valor en la posicion del array de floats 100:",
		array_float[99])

	fmt.Println("Valor en la posicion del array de floats 2º 99:",
		array_float2[98])
	fmt.Println("Valor en la posicion del array de floats 2º 200:",
		array_float2[199])

	var slice_float64 []float64

	cuatro := 4
	diez := 10
	//slice_float64 = array_float2[4:10]
	slice_float64 = array_float2[cuatro:diez]

	fmt.Println("================> ", len(slice_float64), " of ", cap(slice_float64))
	fmt.Println(slice_float64[0])

	fmt.Println(array_float2)
	slice_float64[0] = -999
	fmt.Println("Tras asignar", array_float2)
}

func pruebaMaps() {
	// PARA DICCIONARIOS

	type Persona struct {
		nombre string
		edad   uint16
	}

	var emilio_roma, em_rom, juan_roma, rosa_anad,
		ana_losa, miguel_aron Persona

	emilio_roma.nombre = "Emilio Roma"
	emilio_roma.edad = 34

	em_rom.nombre = "Emilio Roma"
	em_rom.edad = 34

	juan_roma.nombre = "Juan Roma"
	juan_roma.edad = 80

	rosa_anad.nombre = "Rosa Anad"
	rosa_anad.edad = 82

	ana_losa.nombre = "Ana Losa"
	ana_losa.edad = 35

	miguel_aron.nombre = "Miguel Aron"
	miguel_aron.edad = 65

	type Padres struct {
		padre, madre Persona
	}

	var padres_emilio_roma, pareja Padres
	padres_emilio_roma.madre = rosa_anad
	padres_emilio_roma.padre = juan_roma

	pareja.madre = ana_losa
	pareja.padre = miguel_aron

	var pareja2 Padres

	pareja2 = pareja

	pareja2.padre.edad = 10
	pareja2.madre.edad = 12

	var diccionario map[string]string

	diccionario = make(map[string]string)
	diccionario["atrabilario"] =
		"Que tiene mal caracter y se irrita con faciliad. Que es raro o extravagante."
	diccionario["miserable"] =
		"Ruín o canalla. Extremadamente tacaño. Extremadamente pobre."

	fmt.Println("Las palabras guardadas actualmente")
	fmt.Println(diccionario)
	delete(diccionario, "atrabilario")
	fmt.Println("Las palabras guardadas actualmente tras borrar atrabilario")
	fmt.Println(diccionario)

	var familia1 map[Persona]Padres

	familia1 = make(map[Persona]Padres)
	familia1[emilio_roma] = padres_emilio_roma
	fmt.Println("Los padres de ", emilio_roma, " son ",
		familia1[emilio_roma])

	var familia2 = map[Persona]Padres{juan_roma: pareja}
	fmt.Println("Los padres de ", juan_roma, " son ",
		familia2[juan_roma], " lo que resulta un poco inquietante.")

	var familia3 = map[Persona]Padres{juan_roma: pareja}
	fmt.Println(familia3)
}

func main() {

	pruebaArraysSlices()
	pruebaMaps()
	ejArraysSlices()
	//ejsMaps()
}
