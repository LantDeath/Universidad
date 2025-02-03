package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	resources "github.com/hajimehoshi/ebiten/v2/examples/resources/images/flappy"
)

const (
	screenWidth  = 1200
	screenHeight = 900
	ModeGame     = "game"
	ModeGameOver = "gameover"
	initXNave    = 542
	initYNave    = 0
	gravity      = 8.5
)

var (
	g           *Game
	gopherImage *ebiten.Image
	treeImage   *ebiten.Image
	img         *ebiten.Image
	err         error
)

func init() {

	rand.Seed(time.Now().UnixNano())

	img, _, err := image.Decode(bytes.NewReader(resources.Gopher_png))
	img, _, err = ebitenutil.NewImageFromFile("ave_bebe.png")
	if err != nil {
		log.Fatal(err)
	}
	gopherImage = ebiten.NewImageFromImage(img)

	img2, _, err2 := ebitenutil.NewImageFromFile("nube.png")
	if err2 != nil {
		log.Fatal(err2)
	}
	treeImage = ebiten.NewImageFromImage(img2)

}

// PARTE PLANTILLA Y PARTE DE LA LÓGICA DEL JUEGO
type Game struct {
	//mode int

	// POSICION DEL PERSONAJE Y EL RESTO DE PARÁMETROS DEL JUEGO
	x_0                     int
	y_0                     int
	current_x               int
	current_y               int
	x16                     int
	y16                     int
	vx_0                    float64
	vy_0                    float64
	current_vx              float64
	current_vy              float64
	time_per_tick           float64
	vlcty_x_left_per_click  float64
	vlcty_x_right_per_click float64
	vlcty_y_up_per_click    float64
	vlcty_y_down_per_click  float64
	scale_gopher            float64
	tree_x                  int
	tree_y                  int
	x_to_pixel              float64
	y_to_pixel              float64
	starTime                time.Time
	elapsed                 time.Duration
}

// PLANTILLA
func NewGame() *Game {
	g = &Game{}
	g.init()
	g.starTime = time.Now()
	return g
}

func (g *Game) init() {

	// SE INICIALIZAN PARÁMETROS DE TODO TIPO DEL JUEGO (PARA LA REPRESENTACIÓN GRÁFICA, PARA LA SIMULACIÓN DEL PROCESO)
	g.scale_gopher = 0.4
	g.x16 = initXNave
	g.y16 = initYNave
	g.x_0 = initXNave
	g.y_0 = initYNave
	g.current_x = g.x_0
	g.current_y = g.y_0
	g.vx_0 = 0.75
	g.vy_0 = 0.0
	g.current_vx = g.vx_0
	g.current_vy = g.vy_0
	g.time_per_tick = (1 / 360.0)
	// g.time_per_tick = 1 / 10.0
	g.vlcty_x_left_per_click = 1.5
	g.vlcty_y_up_per_click = 1.5
	g.vlcty_x_right_per_click = 1.0
	g.vlcty_y_down_per_click = 1.0
	g.x_to_pixel = 350.0
	g.y_to_pixel = 400.0

	// GESTIÓN DE LA IMAGEN DEL ESCENARIO (PLANTILLA)
	tree_w, tree_y := treeImage.Size()
	g.tree_x = rand.Intn(screenWidth - tree_w)
	g.tree_y = (screenHeight / 2) + rand.Intn(screenHeight/2-tree_y)
	//log.Println("========================================> Tree x:", g.tree_x, " Tree y:", g.tree_y)
}

// FUNCIÓN DE UTILIDAD PARA LA SIMULACIÓN DEL PROCESO FÍSICO
func (g *Game) nextPositionVelocity(vx_t_1, vy_t_1, ax_t_1, ay_t_1 float64) (x, y, vx, vy float64) {
	// A PARTIR DE LA POSICIÓN (x,y) ACTUAL Y LA VELOCIDAD ACTUAL (vx,vy)
	// SE CALCULA LA VELOCIDAD ACTUALIZADA Y EL INCREMENTO EN LA POSICIÓN DEBIDO A LA VELOCIDAD Y ACELERACIÓN EN EL TIEMPO TRANSCURRIDO
	//log.Println("===>", "Vx(t-1)", vx_t_1, "Vy(t-1)", vy_t_1, "Ax(t-1)", ax_t_1, "Ay(t-1)", ay_t_1)

	// SE ACTUALIZA VELOCIDAD EN FUNCIÓN DE LA ACELERACIÓN (MUY GENERAL YA QUE EN X NO HAY ES 0)
	vx = vx_t_1 + ax_t_1*g.time_per_tick
	vy = vy_t_1 + ay_t_1*g.time_per_tick

	// x no es uniformemente acelerado x = v * t se calcula el incremento
	// UNA CURIOSIDAD... APROXIMAR EL NÚMERO EN PIXELES DE DELTA X A LA VELOCIDAD ACTUALIZADA TERMINA SIENDO UN MOVIMIENTO NATURAL
	// g.x_to_pixel = 1
	//x = vx_t_1 + (ax_t_1 * g.time_per_tick * g.time_per_tick / 2)
	x = vx * g.time_per_tick
	// y es uniformemente acelerado el incremento en y es vy* t + 1/2(at^2)
	// UNA CURIOSIDAD... APROXIMAR EL NÚMERO EN PIXELES DE DELTA X A LA VELOCIDAD ACTUALIZADA TERMINA SIENDO UN MOVIMIENTO NATURAL
	// g.y_to_pixel = 1
	// y = vy_t_1 + (ay_t_1 * g.time_per_tick * g.time_per_tick / 2)
	y = vy_t_1*g.time_per_tick + (ay_t_1 * g.time_per_tick * g.time_per_tick / 2)
	//log.Println("<===", "x(t)", x, "y(t)", y, "Vx(t)", vx, "Vy(t)", vy)
	return x, y, vx, vy

}

func (g *Game) Update() error {
	var current_ax, current_ay float64
	current_ax = 0
	current_ay = gravity
	g.elapsed = time.Since(g.starTime)

	if !g.GopherInTree() {

		if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
			//current_ay = gravity + g.acc_per_click
			g.current_vy += g.vlcty_y_down_per_click
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
			//current_ay = gravity - g.acc_per_click
			g.current_vy -= g.vlcty_y_up_per_click
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
			//current_ax = g.acc_per_click
			g.current_vx += g.vlcty_x_right_per_click
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
			//current_ax = -g.acc_per_click
			g.current_vx -= g.vlcty_x_left_per_click
		}
		//x, y, vx, vy := g.nextPositionVolcity(g.current_vx, g.current_vy, 0.0, gravity)
		x, y, vx, vy := g.nextPositionVelocity(g.current_vx, g.current_vy, current_ax, current_ay)
		g.current_vx = vx
		g.current_vy = vy

		// g.current_x += int(math.Round(x))
		// g.current_y += int(math.Round(y))
		// PASAR A PIXELES

		limitX := g.current_x + int(math.Round(x*g.x_to_pixel))
		limitY := g.current_y + int(math.Round(y*g.y_to_pixel))

		w2, h2 := gopherImage.Size()
		scaledWidth := int(float64(w2) * g.scale_gopher)
		scaledHeight := int(float64(h2) * g.scale_gopher)

		if limitX < 0 {
			limitX = 0
		}

		if limitX > screenWidth-scaledWidth {
			limitX = screenWidth - scaledWidth
		}

		if limitY < 0 {
			limitY = 0
		}

		if limitY > screenHeight-scaledHeight {
			limitY = screenHeight - scaledHeight
		}

		g.current_x = limitX
		g.current_y = limitY
		g.x16 = g.current_x
		g.y16 = g.current_y

	}

	return nil
}

// FUNCIONES DE UTILIDAD USADAS EN EL DIBUJADO DE LA ESCENA
func (g *Game) drawGopher(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(g.scale_gopher, g.scale_gopher)
	op.GeoM.Translate(float64(g.x16), float64(g.y16))

	screen.DrawImage(gopherImage, op)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Gopher at (%d,%d). Tree (%d,%d)", g.x16, g.y16, g.tree_x, g.tree_y))
}

func (g *Game) drawTree(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w, h := treeImage.Size()
	op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
	op.GeoM.Translate(float64(w)/2.0, float64(h)/2.0)
	op.GeoM.Translate(float64(g.tree_x), float64(g.tree_y))
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(treeImage, op)
}

// PLATILLA AQUÍ SE INVOCAN LAS FUNCIONES DE DIBUJADO
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 173, G: 216, B: 230, A: 255})
	g.drawTree(screen)
	g.drawGopher(screen)
	elapsedSeconds := int(g.elapsed.Seconds())
	timerText := fmt.Sprintf("Tiempo: %02d:%02d", elapsedSeconds/60, elapsedSeconds%60)
	textWidth := len(timerText) * 6
	PositionX := screenWidth - textWidth - 10
	PositionY := 10
	ebitenutil.DebugPrintAt(screen, timerText, PositionX, PositionY)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) GopherInTree() bool {
	// LÓGICA DEL JUEGO, DETERMINACIÓN DE SI SE HA LLEGADO AL ÁRBOL
	w, h := treeImage.Size()
	w2, h2 := gopherImage.Size()
	//log.Println("####################### Gopher at ", g.current_x, "-", g.current_y, " And tree (", w, "-", h, ") at ", g.tree_x, "-", g.tree_y)
	log.Println("####################### Checking ", g.tree_x, "<=", g.current_x+int(float64(w2)*g.scale_gopher*0.5), "<=", g.tree_x+w, " && ",
		g.tree_y, " <= ", g.current_y+int(float64(h2)*g.scale_gopher*0.5), " <= ", g.tree_y+h)
	// if (g.current_x+(w2/4) >= g.tree_x) && (g.current_x+(w2/4) <= g.tree_x+w) && (g.current_y+(h2/4) >= g.tree_y) && (g.current_y+(h2/4) <= g.tree_y+h) {
	// 	return true
	// } else {
	// 	return false
	// }
	if (g.tree_x <= g.current_x+int(float64(w2)*g.scale_gopher*0.5)) && (g.current_x+int(float64(w2)*g.scale_gopher*0.5) <= g.tree_x+w) &&
		(g.tree_y <= g.current_y+int(float64(h2)*g.scale_gopher*0.5)) && (g.current_y+int(float64(h2)*g.scale_gopher*0.5) <= g.tree_y+h) {
		log.Println("¡ATERRIZASTE!")
		log.Println(true)
		return true
	} else {
		return false
	}

}

func main() {

	// PLANTILLA PARA ESTOS PROYECTOS
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Fly baby")
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}

}
