package main

import (
	"fmt"
	// MATEMATICAS "math"
)

func testStrings() {
	var saludo string = "Hola a migos, bienvenidos"
	var personas string = ": Pedro, Luis, Miguel"
	var nombre1, nombre2 string = "Eva", "Laura"
	var saludo_completo string = saludo + personas
	var igualdad bool

	igualdad = nombre1 == nombre2

	// Los strings son inmutables
	// saludo[7]= 'n'
	// ERROR DE COMPILACIÓN

	fmt.Println("Ejemplo de string " + saludo_completo)
	fmt.Println("Muestra de comparación ¿Es ", nombre1, " igual a ", nombre2, " ?: ", igualdad)
}

func testArrays() {
	// DECLARACIÓN
	var nombres [5]string
	// ... CON INICIALIZACIÓN
	var filiacion = [...]string{"ONU", "UNIR", "EU", "CAM", "UPV"}

	// ... TAMBIÉN EN POSICIONES CONCRETAS
	var origen = [...]string{3: "CAM", 4: "UPV"}

	// 	PRUEBA DE COPIA DE UN ARRAY EN OTRO Y DE COMPARACIÓN DE DOS ARRAYS
	var copia_filiacion [5]string

	copia_filiacion = filiacion

	fmt.Println("Tras copiar un array en otro ", copia_filiacion, " ¿son iguales? ", filiacion == copia_filiacion)

	copia_filiacion[1] = "MIT"

	fmt.Println("Tras modificar una de las dos. La original ", filiacion, " y la copia ", copia_filiacion, " ¿son iguales? ", filiacion == copia_filiacion)

	// INDEXACIÓN
	nombres[0] = "Alberto"
	nombres[1] = "Blanca"
	nombres[2] = "Cesar"
	nombres[3] = "Domingo"
	nombres[4] = "Eva"

	// CONTROL DE LÍMITES DE ÍNDICES
	/*
		nombres[5] = "No debería poder asignar un nombre aquí"
		nombres[-1] = "No debería poder asignar un nombre aquí"
	*/

	fmt.Println("Éstas son las personas ", nombres, " Y estos las filiaciones ", filiacion, " Y estos los orígenes ", origen)
	fmt.Println("Más en detalle ")
	fmt.Println(nombres[0], "(", filiacion[0], ") desde ", origen[0])
	fmt.Println(nombres[1], "(", filiacion[1], ") desde ", origen[1])
	fmt.Println(nombres[2], "(", filiacion[2], ") desde ", origen[2])
	fmt.Println(nombres[3], "(", filiacion[3], ") desde ", origen[3])
	fmt.Println(nombres[4], "(", filiacion[4], ") desde ", origen[4])

}

// PARA LAS PRUEBAS DE STRUCTS DEFINIMOS UNOS CUANTOS
// OBS. MAYÚSCULAS PARA EXPORTAR

func testStructs() {

	// DECLARACION CON INICIALIZACIÓN CON STRUCT LITERAL
	var JuanLopez10_198 Persona = Persona{"Juan Lopez", 10, 1.98}
	var LuisaLopez40_165 Persona

	// QUÉ TIENE  UN STRUCT ANTES DE DARLE VALOR: VALORES CERO
	fmt.Println("Luisa antes de darle valores :", LuisaLopez40_165)

	// ACCESO A LAS COMPONENTES
	LuisaLopez40_165.Nombre = "Luis Lopez"
	LuisaLopez40_165.Edad = 4 * JuanLopez10_198.Edad
	LuisaLopez40_165.Estatura = 1.65

	fmt.Println("Se han definido dos personas :", JuanLopez10_198, " y ", LuisaLopez40_165)

	// COMPARACIÓN DE ORDEN
	fmt.Println("¿Es Juan igual ", JuanLopez10_198, " que Luisa", LuisaLopez40_165, " ?:", JuanLopez10_198 == LuisaLopez40_165)
	fmt.Println("Pero Luisa sí es igual a sí misma ", LuisaLopez40_165 == LuisaLopez40_165)

}

type Persona struct {
	Nombre   string
	Edad     int
	Estatura float64
}

func ejsArraysStructs() {
	// 1. DEFINE UN ARRAY DE PERSONAS
	var array_personas [3]Persona
	persona1 := Persona{"Carlos Pascual", 22, 1.81}
	persona2 := Persona{"Lola Ferran", 36, 1.69}
	// 2. DECLARA UNA VARIABLE DE TIPO PERSONA QUE TE REPRESENTE (PUEDES INVENTARTE LOS DATOS)
	me := Persona{"Liz", 99, 2.95}
	// 3. INSERTA LAS TRES PERSONAS EN EL ARRAY (LAS DOS DEL FUENTE ORIGINAL Y A TI MISMO)
	array_personas[0] = me
	array_personas[1] = persona1
	array_personas[2] = persona2
	// 4. IMPRIME EL ARRAY DE PERSONAS
	fmt.Print("Nuestras personas son:", array_personas)

}

func ejStrings() {
	// 1. DEFINE TRES VARIABLES DE TIPO STRING, AL MENOS UNA DE ELLAS INICIALIZADA CON TU NOMBRE Y DOS APELLIDOS (PUEDES INVENTARTELOS)
	var mate1, mate2, me string = "Carlos Pascual Bolaños", "Lola Ferran Hernández", "Liz Núñez Torres"
	// 2. IMPRIME UN TEXTO MOSTRANDO EL VALOR DE LAS TRES VARIABLES
	fmt.Println(me, "estudia con", mate2, "el lenguaje de programación Go y", mate1, "va a ayudarles con sus apuntes.")
}

func main() {

	testStrings()
	testArrays()
	testStructs()

	ejStrings()
	ejsArraysStructs()

}
