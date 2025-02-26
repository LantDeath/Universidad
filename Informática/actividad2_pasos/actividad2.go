package main

import (

	//"math/rand"

	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

//"math/rand"

type AEMET_estaciones struct {
	Altitud    string `json:"altitud"`
	Indicativo string `json:"indicativo"`
	Indsinop   string `json:"indsinop"`
	Latitud    string `json:"latitud"`
	Longitud   string `json:"longitud"`
	Nombre     string `json:"nombre"`
	Provincia  string `json:"provincia"`
}

type AEMET_query_respuesta struct {
	Descripcion string `json:"descripcion"`
	Estado      int    `json:"estado"`
	Datos       string `json:"datos"`
	Metadatos   string `json:"metadatos"`
}

func main() {
	var respuesta_datos_estaciones []AEMET_estaciones

	respuesta_datos_estaciones = getEstacionesDeAEMET()

	writeSquare("Imagen_raiz")
	writeLoge("Imagen_logaritmo")
	writeCosine("Función_Coseno")
	writeHistogram("Histograma")

	writePlotAltitud(respuesta_datos_estaciones)
	PlotLongitud(respuesta_datos_estaciones)
	PlotLatitud(respuesta_datos_estaciones)

	var respuesta_maritima_datos []AEMET_query_maritima

	respuesta_maritima_datos = get_prediccion_maritima()
	graficarVelMedia(respuesta_maritima_datos)

}

func writeSquare(title string) {
	var x = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var y [10]int

	for i := 0; i < len(x); i++ {
		y[i] = x[i] * x[i]
	}

	p := plot.New()
	puntos := make(plotter.XYs, len(x))

	for i := range puntos {
		puntos[i].X = float64(x[i])
		puntos[i].Y = float64(y[i])
	}

	p.Title.Text = "Gráfica de Raíz"
	p.X.Label.Text = "Abcisas"
	p.Y.Label.Text = "Ordenadas"

	_ = plotutil.AddLines(p, "Grafico", puntos)
	_ = p.Save(700, 700, title+".png")
}

func writeLoge(title string) {
	const size = 10
	var x [size]float64
	var y [size]float64

	// Asignar valores a x e y para representar ln(x)
	for i := 1; i <= size; i++ { // Comienza en 1 para evitar ln(0)
		x[i-1] = float64(i)       // Asignar valores a x del 1 al tamaño (10)
		y[i-1] = math.Log(x[i-1]) // Calcular ln(x)
	}

	// Crear el gráfico
	p := plot.New()
	puntos := make(plotter.XYs, len(x))

	// Asignar los puntos al gráfico
	for i := range puntos {
		puntos[i].X = x[i]
		puntos[i].Y = y[i]
	}

	// Configurar el título y etiquetas
	p.Title.Text = "Gráfica Logaritmo Natural"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "ln(X)"

	// Crear una línea personalizada
	line, err := plotter.NewLine(puntos)
	if err != nil {
		panic(err)
	}
	line.Color = plotutil.Color(1) // Rojo para la línea

	// Añadir la línea al gráfico
	p.Add(line)

	// Guardar la gráfica como imagen
	if err := p.Save(700, 700, title+".png"); err != nil {
		panic(err)
	}
}
