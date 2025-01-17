// PROFESOR: JAVIER FERANDEZ COTO CORDOBA.
// ESTUDIANTE: LIZANGEL NÚÑEZ TORRES.
// ACTIVIDAD 1.

package main

import (
	"encoding/json"

	"io/ioutil"
	"log"
	"net/http"

	"fmt"
	"image"
	"os"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/exp/shiny/widget"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

type NASA_apod_response struct {
	Copyright       string `json:"copyright"`
	Date            string `json:"date"`
	Explanation     string `json:"explanation"`
	Hdurl           string `json:"hdurl"`
	Media_type      string `json:"media_type"`
	Service_version string `json:"service_version"`
	Title           string `json:"title"`
	Url             string `json:"url"`
}

func decode(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	m, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("could not decode %s: %v", filename, err)
	}
	return m, nil
}

func show_image() {
	driver.Main(func(s screen.Screen) {
		src, err := decode("2024-11-18-2_planeteclipse_spitzer_2005-04-05.jpg") // Modificamos el nombre de la extensión por la obtenida en el JSON
		if err != nil {
			log.Fatal(err)
		}
		w := widget.NewSheet(widget.NewImage(src, src.Bounds()))

		// Declaramos lo siguiente para obtener las dimensiones de la imagen automáticamente.
		imgBounds := src.Bounds()
		imgWidth := imgBounds.Dx()  // Ancho de la imagen
		imgHeight := imgBounds.Dy() // Altura de la imagen

		if err := widget.RunWindow(s, w, &widget.RunWindowOptions{
			NewWindowOptions: screen.NewWindowOptions{
				Title: "Light From A Distant Planet",

				Width:  imgWidth,
				Height: imgHeight,
			},
		}); err != nil {
			log.Fatal(err)
		}
	})
}

func getImageFromNASA(fecha string) {
	fmt.Println("Se va a pedir la imagen de la fecha " + fecha)

	res1, err1 := http.Get("https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY&date=" + fecha)
	if err1 != nil {
		log.Fatal(err1)
	}
	//body, err := io.ReadAll(res.Body)
	body1, err1 := ioutil.ReadAll(res1.Body)
	res1.Body.Close()
	if res1.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res1.StatusCode, body1)
	}
	if err1 != nil {
		log.Fatal(err1)
	}
	//fmt.Printf("%s \n\n de tipo %s", body1, body1.(type))

	var rspns NASA_apod_response
	//b := []byte(`{"copyright":"date	"2005-04-05" explanation	"Light emitted by a planet far beyond our Solar System has been identified for the first time.  The planet, illustrated in the above drawing, had its light detected by comparing the brightness of only the parent star, when the planet was behind the star, to the light emitted when both the planet and its parent star were visible.   The Earth-trailing Spitzer Space Telescope made the observation in infrared light, where the intrinsic glow of the planet outshines the light it reflects from its central star.  The direct observation of light allowed a measurement of both the temperature and size of the planet: HD 209458b.  Planet HD 209458b was confirmed to be larger than expected for its mass and on an orbit around its parent star that was unexpectedly close to a circle."hdurl	"https://apod.nasa.gov/ap…teclipse_spitzer_big.jpg"media_type	"image"service_version	"v1"title	"Light From A Distant Planet"url "https://apod.nasa.gov/ap…laneteclipse_spitzer.jpg"
	// err2 := json.Unmarshal(b, &rspns)
	err2 := json.Unmarshal(body1, &rspns)
	if err2 != nil {
		log.Fatal(err2)
	}
	if rspns.Media_type != "image" {
		fmt.Printf("El tipo de recurso %s no es una imagen", rspns.Media_type)
		os.Exit(0)
	}

	//res, err := http.Get("https://apod.nasa.gov/ap…laneteclipse_spitzer.jpg")
	fmt.Printf("Obtenemos la imagen de la nasa %s", rspns.Hdurl)
	fmt.Printf("\nQue tiene como título %s", rspns.Title)
	// res, err := http.Get(rspns.Hdurl)
	res, err := http.Get(rspns.Url)
	if err != nil {
		log.Fatal(err)
	}
	//body, err := io.ReadAll(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s", body)
	ioutil.WriteFile("2024-11-18-2_planeteclipse_spitzer_2005-04-05.jpg", body, 0666) // Ponemos la extensión del fichero.
	log.SetFlags(0)
}

func main() {
	//res, err := http.Get("http://www.google.com/robots.txt")

	var fecha string

	if len(os.Args) < 2 {
		fecha = "2005-04-05" // Escogemos una fecha, en mi caso he puesto "2005-04-05" donde la imagen se llama "Light From A Distant Planet".
		//fecha = "1967-08-25"
		//fecha = "1995-08-25"

	} else {
		fecha = os.Args[1]
	}

	getImageFromNASA(fecha)

	show_image()

}
