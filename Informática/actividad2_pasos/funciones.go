package main

import (
	"bytes"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"encoding/json"

	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"strconv"

	"math"
)

func getAEMETQuerySimple(query string) []byte {
	var err error
	var res *http.Response
	var body []byte

	res, err = http.Get(query)

	if err != nil {
		log.Fatal(err)
	}

	body, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func getEstacionesDeAEMET() []AEMET_estaciones {
	var body []byte
	var respuesta_query_estaciones AEMET_query_respuesta
	var respuesta_datos_estaciones []AEMET_estaciones

	AEMET_clave_temporal := "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJsaXphbmdlbG51bmV6QGdtYWlsLmNvbSIsImp0aSI6ImVjYzBkNTk0LWI5MTktNDU0OC1iY2I2LTVkMTE4Njk0NzI5ZiIsImlzcyI6IkFFTUVUIiwiaWF0IjoxNzM2MDk0NzE2LCJ1c2VySWQiOiJlY2MwZDU5NC1iOTE5LTQ1NDgtYmNiNi01ZDExODY5NDcyOWYiLCJyb2xlIjoiIn0.wbZBsSxKVHKozEiIxiM4gTj11zCKzaVXnl-KJEd8JfY"
	AEMET_query_estaciones := "https://opendata.aemet.es/opendata/api/valores/climatologicos/inventarioestaciones/todasestaciones?api_key=" + AEMET_clave_temporal

	body = getAEMETQuerySimple(AEMET_query_estaciones)
	fmt.Println("=======================>>>>> Body of AEMET's response ", bytes.NewBuffer(body).String())
	_ = json.Unmarshal(body, &respuesta_query_estaciones)

	body = getAEMETQuerySimple(respuesta_query_estaciones.Datos)

	// PLANTILLA, RELLENA UNA COLECCIÓN EN ESTE CASO UN ARRAY DECLARADO ADECUADAMENTE.
	_ = json.Unmarshal(body, &respuesta_datos_estaciones)

	return respuesta_datos_estaciones

}

func writePlotAltitud(respuesta_datos_estaciones []AEMET_estaciones) {
	puntos := make(plotter.XYs, len(respuesta_datos_estaciones))

	for i := range puntos {
		puntos[i].X = float64(i)
		puntos[i].Y, _ = strconv.ParseFloat(respuesta_datos_estaciones[i].Altitud, 64)
	}

	p := plot.New()

	p.Title.Text = "Example"
	p.X.Label.Text = "Abcisas"
	p.Y.Label.Text = "Ordenadas"

	linePoints, err := plotter.NewLine(puntos)
	if err != nil {
		panic(err)
	}

	linePoints.Color = plotutil.Color(5)

	p.Add(linePoints)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "Altitud.png"); err != nil {
		panic(err)
	}
}

//                                             --------------------> Mostrar gráficas de diferentes funciones <--------------------

func writeCosine(title string) {
	const size = 100    // Damos el el tamaño de la constante "size" que consta de 99 posiciones y 100 ubicaciones.
	var x [size]float64 // Declaramos el array para x donde su tamaño es "size" y cada elemento del índice size es un float64.
	var y [size]float64 // Declaramos el array para y donde su tamaño es "size" y cada elemento del índice size es un float64.

	// Hacer la gráfica de los puntos que pasan por X e Y por medio de un bucle.
	for i := 0; i < size; i++ {
		/*
			"i := 0" --> Inicia desde 0 y termine hasta el tamaño establecido de la constate "size".
			"i < size" --> El bucle se ejecutara siempre que "i" sea menor que "size", cuando "i" llegue a "size" el bucle terminará.
			"i++" Aumenta el valor de "i" de 1 en 1, "i++" es la forma compacta de i = i + 1
		*/
		x[i] = float64(i) * (2 * math.Pi / float64(size)) // En esta línea creamos el rango por donde pasaran todas la ubiaciones del array y dividimos el rango por el número total de ubicaciones para que el recorrido este bien distribuido.
		y[i] = math.Cos(x[i])                             // "y[i]" albergará el coseno de los puntos "x[i]"

		//Creamos los puntos d ela gráfica:
		p := plot.New()                     // Creamos una instancia de un gráfico que se almacenara en la variable "p"
		puntos := make(plotter.XYs, len(x)) // Hacemos los puntos 2D con de X e Y con "plot.XYs", y el "make" crea un slice desde plotter del tamaño de los valores de X.

		// A continuación, usaremos un bucle para asignar las coordenadas correspondientes.
		for i := range puntos {
			puntos[i].X = x[i]
			puntos[i].Y = y[i]
		}

		// Añadimos las leyendas de la gráfica:
		p.Title.Text = "Gráfica Coseno"
		p.X.Label.Text = "Radianes"
		p.Y.Label.Text = "Coseno"

		line, err := plotter.NewLine(puntos)
		if err != nil {
			panic(err)
		}

		line.Color = plotutil.Color(2)

		p.Add(line)

		// Guardamos y termianos la imagen:
		_ = p.Save(700, 700, title+".png")

	}

}

//                                             --------------------> Mostrar otras gráficas distintas del diagrama de líneas que unen puntos <--------------------

func writeHistogram(title string) {
	// Generamos datos aleatorios con el paquete  "math/rand"
	rand.Seed(time.Now().UnixNano())
	data := make(plotter.Values, 100)
	for i := range data {
		data[i] = rand.NormFloat64() //Genera los números
	}

	// Hacemos el nuevo plot.
	p := plot.New()
	p.Title.Text = "Histograma"
	p.X.Label.Text = "Valores"
	p.Y.Label.Text = "frecuencia"

	// Crear el histograma
	h, err := plotter.NewHist(data, 16)
	if err != nil {
		log.Fatal("Error en Histograma", err)
	}

	// Personalizar el histograma
	h.FillColor = color.RGBA{B: 100, G: 200, R: 150, A: 125} // Color de las barras
	p.Add(h)

	// Guardar gráfica
	if err := p.Save(700, 700, title+".png"); err != nil {
		log.Fatal("Error en histograma", err)
	}

}

//                                             --------------------> Función longitud (Código del pdf) <--------------------

func longitudAFloat(longitud string) float64 {
	aux_return, _ := strconv.ParseFloat(longitud[0:len(longitud)-1], 64)
	if longitud[len(longitud)-1] == 'E' {
		return aux_return
	} else {
		return -aux_return
	}
}

//                                             --------------------> Función latitud (Código del pdf) <--------------------

func latitudAFloat(latitud string) float64 {
	aux_return, _ := strconv.ParseFloat(latitud[0:len(latitud)-1], 64)
	if latitud[len(latitud)-1] == 'N' {
		return aux_return
	} else {
		return -aux_return
	}
}

//                                             --------------------> Función Para graficar latitud <--------------------

func PlotLatitud(respuesta_datos_estaciones []AEMET_estaciones) {
	puntos := make(plotter.XYs, len(respuesta_datos_estaciones))

	for i := range puntos {
		puntos[i].X = float64(i)
		puntos[i].Y = latitudAFloat(respuesta_datos_estaciones[i].Latitud)
	}

	p := plot.New()

	p.Title.Text = "Latitud de las estaciones"
	p.X.Label.Text = "Estación"
	p.Y.Label.Text = "Latitud"

	// Generamos la gráfica con una imagen de línea más el color
	line, err := plotter.NewLine(puntos)
	if err != nil {
		panic(err)
	}
	line.Color = plotutil.Color(1)

	p.Add(line)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "Estaciones_lat.png"); err != nil {
		panic(err)
	}
}

//                                             --------------------> Función Para graficar longitud <--------------------

func PlotLongitud(respuesta_datos_estaciones []AEMET_estaciones) {
	puntos := make(plotter.XYs, len(respuesta_datos_estaciones))

	for i := range puntos {
		puntos[i].X = float64(i)
		puntos[i].Y = longitudAFloat(respuesta_datos_estaciones[i].Longitud)
	}

	p := plot.New()

	p.Title.Text = "Longitud de las estaciones"
	p.X.Label.Text = "Estación"
	p.Y.Label.Text = "Longitud"

	line, err := plotter.NewLine(puntos)
	if err != nil {
		panic(err)
	}
	line.Color = plotutil.Color(4)

	p.Add(line)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "Estaciones_long.png"); err != nil {
		panic(err)
	}

}
