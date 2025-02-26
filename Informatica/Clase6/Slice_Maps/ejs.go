package main

import (
	"fmt"
)

func ejArraysSlices() {
	fmt.Println("Ejs slices")
	// HACER LA SEMANA QUE VIENE
	// 1. Definir un array con datos (de lo que queráis)
	var cientificos = [...]string{"Bohr", "Schrödinger", "Heisenberg", "Ramón y Cajal", "Galileo", "Newton", "Farady", "Maxwell", "Curi", "Meitner", "Feyman"}
	fmt.Println("Los cientificos", cientificos, "son unos de los muchos que han contribuido al conocimiento.")

	// 2. Definir un slice y asociarlo con el array anterior desde la posición central hasta la posción a 3/4 partes ( aprox. )
	var cuanticos = cientificos[0:4]
	fmt.Println("Los cientificos", cuanticos, "son de los pioneros en la física de cuantos.")
	var clasicos = cientificos[4:6]
	fmt.Println("Los físicos", clasicos, "fueron quienes pusieron las bases de la física clásica.")

	// 3. Modificar una posición central del slice
	cientificos[0], cientificos[1], cientificos[3] = "Fermi", "Hertz", "Dirac"

	// 4. Mostrar esa posición desde el array y ver que ha cambiado
	fmt.Println("Los nuevos físicos que se unen son:", cientificos[0:4])
}

/*
func ejsMaps() {
	fmt.Println("Ejs Maps")
	// SÓLO LOS AVANZADOS
	// 1. Definir un mapa que tenga como clave un tipo de dato propio (declarado por el programado) el tipo de los valores es el que queráis
	// 2. Asignar dos valores a dos claves
	// 3. Mostrarlo
	// 4. Quitar uno
	// 5. Mostrarlo
}
*/
