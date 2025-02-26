package main

import (
	"fmt"
)

type Prenda struct {
	tipo   string
	talla  int
	precio float64
}

func main() {
	var pantalon Prenda
	var camiseta Prenda
	var doble *Prenda

	pantalon.tipo = "Pantalón"
	pantalon.talla = 42
	pantalon.precio = 45.52

	camiseta.tipo = "Camiseta"
	camiseta.talla = 50
	camiseta.precio = 60

	fmt.Println("El", pantalon.tipo, "que necesito debe ser de talla", pantalon.talla, "aunque sea tan caro que cueste", pantalon.precio, ".")

	fmt.Println("La", camiseta.tipo, "cuesta", camiseta.precio, "y aunque es cara la compraré porque su talla", camiseta.talla, "es única.")

	doble = &pantalon // Ahora doble apunta a pantalon
	doble.precio *= 2 // Se modifica el precio del pantalón a través del puntero

	doble = &camiseta // Ahora doble apunta a camiseta
	doble.precio *= 2 // Se modifica el precio de la camiseta a través del puntero

	fmt.Println("El doble de", pantalon.tipo, "y", camiseta.tipo, "es", pantalon.precio, "y", camiseta.precio)

	// para hacer llamado usar el debug primero y dejar un espacio al final del código para que se pueda ejecutar el código

}
