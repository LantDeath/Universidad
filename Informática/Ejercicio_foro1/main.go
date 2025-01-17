package main

import "fmt"

var numero int = 11
var p_numero *int = &numero

func main() {
    fmt.Println("¡Hola mundo!", numero, *p_numero)
    *p_numero = 12
    fmt.Println("¡Hola mundo!", numero)
    numero = 13
    fmt.Println("¡Hola mundo!", numero)
    fmt.Println("¡Hola mundo!", numero, *p_numero)
}
