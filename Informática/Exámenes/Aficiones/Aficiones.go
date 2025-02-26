package main

import (
	"fmt"
)

// Creamos el struct
type Aficiones struct {
	Nombre         string
	Federado       bool
	HorasSemanales float32
}

func main() {
	// Creamos dos variables para el struct
	var Aficion1 = Aficiones{"Tiro con arco", true, 6}
	var Aficion2 = Aficiones{"Ajedrez", true, 14}

	// Mostramos por pantalla las aficiones
	fmt.Println("Las aficiones son:")
	fmt.Println(Aficion1)
	fmt.Println(Aficion2)

	// Creamos un puntero pAficion
	var pAficion *Aficiones

	// Modificamos "Federado" por medio del puntero
	pAficion = &Aficion1
	pAficion.Federado = false

	pAficion = &Aficion2
	pAficion.Federado = false

	// Hacemos el llamado por pantalla

	fmt.Println("Las aficiones actualizadas son:")
	fmt.Println(Aficion1)
	fmt.Println(Aficion2)
}
