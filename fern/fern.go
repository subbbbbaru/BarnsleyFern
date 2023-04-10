package fern

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
)

var barnsley = [][]float32{{0, 0, 0, 0.16, 0, 0, 0.01},
	{0.85, 0.04, -0.04, 0.85, 0, 1.6, 0.85},
	{0.2, -0.26, 0.23, 0.22, 0, 1.6, 0.07},
	{-0.15, 0.28, 0.26, 0.24, 0, 0.44, 0.07}}

var cyclosorus = [][]float32{{0, 0, 0, 0.25, 0, -0.4, 0.02},
	{0.95, 0.005, -0.005, 0.93, -0.002, 0.5, 0.84},
	{0.035, -0.2, 0.16, 0.04, -0.09, 0.02, 0.07},
	{-0.04, 0.2, 0.16, 0.04, 0.083, 0.12, 0.07}}

var modified = [][]float32{{0, 0, 0, 0.2, 0, -0.12, 0.01},
	{0.845, 0.035, -0.035, 0.82, 0, 1.6, 0.85},
	{0.2, -0.31, 0.255, 0.245, 0, 0.29, 0.07},
	{-0.15, 0.24, 0.25, 0.20, 0, 0.68, .07}}

var culcita = [][]float32{{0, 0, 0, 0.25, 0, -0.14, 0.02},
	{0.85, 0.02, -0.02, 0.83, 0, 1, 0.84},
	{0.09, -0.28, 0.3, 0.11, 0, 0.6, 0.07},
	{-0.09, 0.28, 0.3, 0.09, 0, 0.7, 0.07}}

var fishbone = [][]float32{{0, 0, 0, 0.25, 0, -0.4, 0.02},
	{0.95, 0.002, -0.002, 0.93, -0.002, 0.5, 0.84},
	{0.035, -0.11, 0.27, 0.01, -0.05, 0.005, 0.07},
	{-0.04, 0.11, 0.27, 0.01, 0.047, 0.06, 0.07}}

var tree = [][]float32{{0, 0, 0, 0.5, 0, 0, 0.05},
	{0.42, -0.42, 0.42, 0.42, 0, 0.2, 0.4},
	{0.42, 0.42, -0.42, 0.42, 0, 0.2, 0.4},
	{0.1, 0, 0, 0.1, 0, 0.2, 0.15}}

var bee = [][]float32{{0.6178, 0, 0, -.6178, 0, 1, 0.5},
	{0, -0.786, 0.786, 0, 0.786, 0, 0.5},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0}}

var ferns = [][][]float32{barnsley, cyclosorus, modified, culcita, fishbone, tree, bee}

var (
	xMinFrac float32 = 0
	xMaxFrac float32 = 0
	yMinFrac float32 = 0
	yMaxFrac float32 = 0
	WIDTH            = 1920
	HEIGHT           = 1080
	probs            = make([]float32, 5)
	affines          = make([][]float32, 4)
	fractal          = make([][]float32, 1)
	POINTS           = 100000
	SCALE            = 1
	points           = [][]int{}
)

func MakeFractal() {

	fractal[0] = append(fractal[0], float32(0))

	makeMatrices()
	var options = len(probs) - 1
	for i := 1; i < POINTS; i++ {
		r := rand.Intn(101)
		for j := 0; j < options; j++ {
			if int(probs[j]*100) <= r && r < int(probs[j+1]*100) { // ERRORRRRR!
				transform(affines[j])
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
}

func makeMatrices() {

	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			affines[i] = append(affines[i], barnsley[i][j]) //barnsley[i][j]
		}
		probs[i+1] = probs[i] + barnsley[i][6]
	}
}

func transform(matrix []float32) {
	length := len(fractal)
	// fractal[0] = append(fractal[0], 0)

	// fractal[0] = make([]float32, 1)
	for i := 0; i < 3; i++ {
		fractal = append(fractal[0:], append(make([]float32, 1), float32(0)))
	}
	if fractal[length] == nil {
		fractal = append(fractal[length:], append(make([]float32, 1), float32(0)))
		fractal[length] = append(fractal[length], (matrix[0]*fractal[length-1][0])+(matrix[1]*fractal[length-1][1])+(matrix[4]), (matrix[2]*fractal[length-1][0])+(matrix[3]*fractal[length-1][1])+(matrix[5]))
	} else {
		fractal[length] = append(fractal[length], (matrix[0]*fractal[length-1][0])+(matrix[1]*fractal[length-1][1])+(matrix[4]), (matrix[2]*fractal[length-1][0])+(matrix[3]*fractal[length-1][1])+(matrix[5]))
	}

	fractal = append(fractal, fractal...)
	fractal[0] = append(fractal[0], 0)
	fractal[1] = append(fractal[1], 1)

}

func GeneratePoints() {
	fwidth := xMaxFrac - xMinFrac
	fheight := yMaxFrac - yMinFrac
	xratio := WIDTH / int(fwidth)
	yratio := HEIGHT / int(fheight)
	xmid := xMinFrac + fwidth/2
	factor := 0
	m := 0

	coords := [][]float32{}

	if xratio < yratio {
		factor = xratio
	} else {
		factor = yratio
	}

	for k := 0; k < POINTS; k++ {
		x := math.Round((float64(fractal[k][0])-float64(xmid))*float64(factor)*float64(SCALE) + float64(WIDTH)/2)
		y := math.Round((float64(fractal[k][1]) - float64(yMinFrac)) * float64(factor) * float64(SCALE))
		if 0 <= x && x < float64(WIDTH) && 0 <= y && y < float64(HEIGHT) {
			if coords[int(x)][int(y)] == 0 {
				coords[int(x)][int(y)] = 1
			} else {
				coords[int(x)][int(y)]++
			}
		}
	}
	for i := 0; i < len(coords); i++ {
		k := len(coords[i])
		if k == 0 {
			continue
		}
		for l := 0; l < k; l++ {
			if coords[i][l] == 1 {
				points[m] = append(points[m], i, l)
				m++
			}
		}
	}
	pngg()

}

func pngg() {

	upLeft := image.Point{0, 0}
	lowRight := image.Point{WIDTH, HEIGHT}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	cyan := color.RGBA{100, 200, 200, 0xff}

	for i := 0; i < POINTS; i++ {

		img.SetRGBA(points[i][0], HEIGHT-points[i][1], cyan)

	}
	f, _ := os.Create("image1.png")
	png.Encode(f, img)
	//println(x, y)
}