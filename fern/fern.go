package fern

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"strconv"
)

type Ferns struct {
	Fern       [][]float64
	PointS     int
	Scale      float64
	PlotHeight int
	PlotWidth  int
	gg         int
}

var Barnsley = [][]float64{{0, 0, 0, 0.16, 0, 0, 0.01},
	{0.85, 0.04, -0.04, 0.85, 0, 1.6, 0.85},
	{0.2, -0.26, 0.23, 0.22, 0, 1.6, 0.07},
	{-0.15, 0.28, 0.26, 0.24, 0, 0.44, 0.07}}

var Cyclosorus = [][]float64{{0, 0, 0, 0.25, 0, -0.4, 0.02},
	{0.95, 0.005, -0.005, 0.93, -0.002, 0.5, 0.84},
	{0.035, -0.2, 0.16, 0.04, -0.09, 0.02, 0.07},
	{-0.04, 0.2, 0.16, 0.04, 0.083, 0.12, 0.07}}

var Modified = [][]float64{{0, 0, 0, 0.2, 0, -0.12, 0.01},
	{0.845, 0.035, -0.035, 0.82, 0, 1.6, 0.85},
	{0.2, -0.31, 0.255, 0.245, 0, 0.29, 0.07},
	{-0.15, 0.24, 0.25, 0.20, 0, 0.68, .07}}

var Culcita = [][]float64{{0, 0, 0, 0.25, 0, -0.14, 0.02},
	{0.85, 0.02, -0.02, 0.83, 0, 1, 0.84},
	{0.09, -0.28, 0.3, 0.11, 0, 0.6, 0.07},
	{-0.09, 0.28, 0.3, 0.09, 0, 0.7, 0.07}}

var Fishbone = [][]float64{{0, 0, 0, 0.25, 0, -0.4, 0.02},
	{0.95, 0.002, -0.002, 0.93, -0.002, 0.5, 0.84},
	{0.035, -0.11, 0.27, 0.01, -0.05, 0.005, 0.07},
	{-0.04, 0.11, 0.27, 0.01, 0.047, 0.06, 0.07}}

var Tree = [][]float64{{0, 0, 0, 0.5, 0, 0, 0.05},
	{0.42, -0.42, 0.42, 0.42, 0, 0.2, 0.4},
	{0.42, 0.42, -0.42, 0.42, 0, 0.2, 0.4},
	{0.1, 0, 0, 0.1, 0, 0.2, 0.15}}

var Bee = [][]float64{{0.6178, 0, 0, -.6178, 0, 1, 0.5},
	{0, -0.786, 0.786, 0, 0.786, 0, 0.5},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0}}

const (
	PlotWidth  = 1920 // размер картинки по ширине
	PlotHeight = 1080 // размер картинки по высоте
)

var (
	xMinFrac float64 = 0 // габариты фракталов
	xMaxFrac float64 = 0 // габариты фракталов
	yMinFrac float64 = 0 // габариты фракталов
	yMaxFrac float64 = 0 // габариты фракталов
	// probs            = make([]float64, 5)
	// affines = [][]float64{} //make([][]float64, 4)
	//lenFractal         = 0
	Points = 100000
	radius = 0.69 // для SVG(поле Circle)
	// fractal            = [][]float64{}
	// points = [][]int{}
)

func (fern Ferns) initFractal() [][]float64 {
	fractal := make([][]float64, fern.PointS)
	for i := range fractal {
		fractal[i] = make([]float64, 3)
	}
	return fractal
}
func (fern Ferns) initPoints() [][]int {
	points := make([][]int, fern.PointS)
	for i := range points {
		points[i] = make([]int, 2)
	}
	return points
}

func (fern *Ferns) makeFractal(lenFractal *int) [][]float64 {

	fractal := fern.initFractal()

	probs, affines := fern.makeMatrices()
	var options = len(probs) - 1
	for i := 1; i < fern.PointS; i++ {
		r := rand.Intn(101)
		for j := 0; j < options; j++ {
			if int(probs[j]*100) <= r && r < int(probs[j+1]*100) { // умножаю на 100 чтобы была не нулевая целая часть
				transform(affines[j], fractal, lenFractal)
				fractal[i][2] = probs[j+1]
				break
			}
		}
		if fractal[i][0] < xMinFrac {
			xMinFrac = fractal[i][0]
		}
		if fractal[i][1] < yMinFrac {
			yMinFrac = fractal[i][1]
		}
		if fractal[i][0] > xMaxFrac {
			xMaxFrac = fractal[i][0]
		}
		if fractal[i][1] > yMaxFrac {
			yMaxFrac = fractal[i][1]
		}
	}
	return fractal
}

func (fern *Ferns) makeMatrices() ([]float64, [][]float64) {
	probs := make([]float64, 5)
	affines := make([][]float64, 4)

	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			affines[i] = append(affines[i], fern.Fern[i][j])
		}
		probs[i+1] = probs[i] + fern.Fern[i][6]
	}
	return probs, affines
}

func transform(matrix []float64, fractal [][]float64, lenFractal *int) {
	*lenFractal += 1
	length := *lenFractal
	fractal[length][0] = (matrix[0] * fractal[length-1][0]) + (matrix[1] * fractal[length-1][1]) + (matrix[4])
	fractal[length][1] = (matrix[2] * fractal[length-1][0]) + (matrix[3] * fractal[length-1][1]) + (matrix[5])
}

func (fern *Ferns) generatePoints() [][]int {
	lenFractal := 0
	fractal := fern.makeFractal(&lenFractal)

	fwidth := xMaxFrac - xMinFrac
	fheight := yMaxFrac - yMinFrac
	if int(fwidth) == 0 {
		fwidth = 1
	}
	if int(fheight) == 0 {
		fheight = 1
	}
	xratio := fern.PlotWidth / int(fwidth)
	yratio := fern.PlotHeight / int(fheight)
	xmid := xMinFrac + fwidth/2
	factor := 0
	m := 0

	coords := make([][]float32, fern.PlotWidth)
	for i := range coords {
		coords[i] = make([]float32, fern.PlotHeight)
	}

	if xratio < yratio {
		factor = xratio
	} else {
		factor = yratio
	}

	for k := 0; k < fern.PointS; k++ {
		x := math.Round((fractal[k][0]-xmid)*float64(factor)*fern.Scale + float64(fern.PlotWidth)/2)
		y := math.Round((fractal[k][1] - yMinFrac) * float64(factor) * fern.Scale)
		if 0 <= x && x < float64(fern.PlotWidth) && 0 <= y && y < float64(fern.PlotHeight) {
			if coords[int(x)][int(y)] == 0 {

				coords[int(x)][int(y)] = 1
			} else {
				coords[int(x)][int(y)]++
			}
		}
	}
	points := fern.initPoints()
	for i := 0; i < len(coords); i++ {
		k := len(coords[i])
		if k == 0 {
			continue
		}
		for l := 0; l < k; l++ {
			if coords[i][l] == 1 {
				points[m][0] = i
				points[m][1] = l
				m++
			}
		}
	}
	return points
}

func (fern *Ferns) GenPng() {
	points := fern.generatePoints()

	upLeft := image.Point{0, 0}
	lowRight := image.Point{fern.PlotWidth, fern.PlotHeight}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	cyan := color.RGBA{100, 200, 200, 0xff}

	for i := 0; i < fern.PointS; i++ {

		img.SetRGBA(points[i][0], fern.PlotHeight-points[i][1], cyan)

	}
	r := strconv.Itoa(rand.Intn(101))
	f, _ := os.Create(r + ".png")
	defer f.Close()
	png.Encode(f, img)
}

func (fern *Ferns) GenSvg() {
	points := fern.generatePoints()

	f, _ := os.Create("fern.svg")
	defer f.Close()
	fmt.Fprintf(f, "<svg viewBox='%d 0 %d %d' xmlns='http://www.w3.org/2000/svg'>\n", 0, fern.PlotHeight, fern.PlotHeight)

	for i := 1; i < fern.PointS; i++ {
		fmt.Fprintf(f, "<circle cx='%v' cy='%g' r='%g' fill='cyan'/>\n", float32(points[i][0]), float32(PlotHeight-points[i][1]), radius)
	}
	fmt.Fprint(f, "</svg>")
}
