package main

import (
	"fmt"
)

// Creamos el struct
type Corredor struct {
	Nombre   string
	Apellido string
	Edad     int
	Dorsal   int
	Puesto   int
}

func main() {

	// Definimos la variables  para representar los corredores (dejamos Puesto 0 en todas las variables dado que se inicia la carrera y porque la variable debe tener definidos todos los componentes del struct)
	corredor1 := Corredor{"Juan", "López", 23, 1, 0}
	corredor2 := Corredor{"Abel", "Mandón", 73, 2, 0}
	corredor3 := Corredor{"Luisa", "Risa", 13, 3, 0}
	corredor4 := Corredor{"Ben", "Jonson", 123, 4, 0}

	// Mostramos los corredores:
	fmt.Println("Los corredores son:")
	fmt.Println(corredor1)
	fmt.Println(corredor2)
	fmt.Println(corredor3)
	fmt.Println(corredor4)

	// Creamos un puntero
	var pCorredor *Corredor

	// Mostramos por medio del puntero el resultado de la carrera:

	pCorredor = &corredor1
	pCorredor.Puesto = 4

	pCorredor = &corredor2
	pCorredor.Puesto = 1

	pCorredor = &corredor3
	pCorredor.Puesto = 3

	pCorredor = &corredor4
	pCorredor.Puesto = 2

	// Mostramos por pantalla los posiciones logradas por los corredores:
	fmt.Println("Terminada la carrera, los puesto de los corredores son:")
	fmt.Println(corredor1)
	fmt.Println(corredor2)
	fmt.Println(corredor3)
	fmt.Println(corredor4)

}
