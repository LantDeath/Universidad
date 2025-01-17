package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

/*
------------------------------> Vamos a hacer en este archivo otra consulta a la AEMET.
------------------------------> Además de replicar la actividad obteniendo otros datos.
*/

type AEMET_prediccion_maritima struct {
	Descripcion string `json:"descripcion"`
	Estado      int    `json:"estado"`
	Datos       string `json:"datos"`
	Metadatos   string `json:"metadatos"`
}

type AEMET_query_maritima struct {
	Fecha    string   `json:"fecha"`    // Fecha del día
	TMax     *float64 `json:"tmax"`     // Temperatura máxima (°C)
	TMin     *float64 `json:"tmin"`     // Temperatura mínima (°C)
	Prec     *float64 `json:"prec"`     // Precipitación diaria (mm)
	VelMedia *float64 `json:"velmedia"` // Velocidad media del viento (m/s)
}

func GetAEMETMaritima(query string) []byte {
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

func get_prediccion_maritima() []AEMET_query_maritima {
	var body []byte
	var respuesta_prediccion_maritima AEMET_prediccion_maritima
	var respuesta_datos_maritimos []AEMET_query_maritima

	AEMET_clave_temporal := "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJsaXphbmdlbG51bmV6QGdtYWlsLmNvbSIsImp0aSI6ImVjYzBkNTk0LWI5MTktNDU0OC1iY2I2LTVkMTE4Njk0NzI5ZiIsImlzcyI6IkFFTUVUIiwiaWF0IjoxNzM2MDk0NzE2LCJ1c2VySWQiOiJlY2MwZDU5NC1iOTE5LTQ1NDgtYmNiNi01ZDExODY5NDcyOWYiLCJyb2xlIjoiIn0.wbZBsSxKVHKozEiIxiM4gTj11zCKzaVXnl-KJEd8JfY"
	AEMET_query_prediccion_maritima := "https://opendata.aemet.es/opendata/api/prediccion/maritima/altamar/area/1?api_key=" + AEMET_clave_temporal

	body = GetAEMETMaritima(AEMET_query_prediccion_maritima)
	fmt.Println("=======================>>>>> Body of AEMET's response ", bytes.NewBuffer(body).String())
	_ = json.Unmarshal(body, &respuesta_prediccion_maritima)

	body = GetAEMETMaritima(respuesta_prediccion_maritima.Datos)

	_ = json.Unmarshal(body, &respuesta_datos_maritimos)

	return respuesta_datos_maritimos

}

func graficarVelMedia(respuesta_datos_maritimos []AEMET_query_maritima) {
	// Crear puntos para el gráfico
	puntos := make(plotter.XYs, 0)

	for i, dato := range respuesta_datos_maritimos {
		if dato.VelMedia != nil { // Ignorar valores nulos
			puntos = append(puntos, plotter.XY{X: float64(i), Y: *dato.VelMedia})
		} else {
			fmt.Printf("Dato omitido: Día %d no tiene VelMedia\n", i)
		}
	}

	if len(puntos) == 0 {
		fmt.Println("No hay puntos válidos para VelMedia. Gráfico no generado.")
		return
	}

	p := plot.New()
	p.Title.Text = "Velocidad Media del Viento (VelMedia)"
	p.X.Label.Text = "Índice (Días)"
	p.Y.Label.Text = "Velocidad Media (m/s)"

	linea, err := plotter.NewLine(puntos)
	if err != nil {
		log.Fatalf("Error creando línea de datos: %v\n", err)
	}
	p.Add(linea)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, "VelMedia.png"); err != nil {
		log.Fatalf("Error guardando gráfico: %v\n", err)
	}

	fmt.Println("VelMedia.png")
}
