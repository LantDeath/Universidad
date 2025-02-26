package main

import (
	"fmt"
)

type Prenda struct {
	TipoDePrenda string
	Talla        int
	Precio       float64
}

func main() {

	var camisa = Prenda{
		TipoDePrenda: "camisa",
		Talla:        33,
		Precio:       20.50,
	}

	var abrigo = Prenda{
		TipoDePrenda: "Abrigo",
		Talla:        36,
		Precio:       60.75,
	}

	fmt.Println("Las prendas son: ")
	fmt.Println(camisa)
	fmt.Println(abrigo)

	// creamos los punteros:

	punteroCamisa := &camisa
	punteroAbrigo := &abrigo

	//accedemos al precio por medio de los punteros y duplicamos:

	punteroCamisa.Precio *= 2
	punteroAbrigo.Precio *= 2

	fmt.Println("Las prendas con precio duplicado son:")
	fmt.Println(camisa)
	fmt.Println(abrigo)
}
