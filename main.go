package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
)

const (
	width  = 1920
	height = 1080
	xMIN   = -2.182
	xMAX   = 2.6558
	yMIN   = 0
	yMAX   = 9.9983

	radius = 0.69

	//width, height = 600, 320            // canvas size in pixels
	cells   = 100                 // number of grid cells
	xyrange = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale = width / 2 / xyrange // pixels per x or y unit
	zscale  = height * 0.4        // pixels per z unit
	angle   = math.Pi / 6         // angle of x, y axes (=30°)
)

func xPointInt(x float32) int {
	return int((x - xMIN) * (width / 2) / (xMAX - xMIN))
}
func yPointInt(y float32) int {
	return int((y - yMIN) * (height / 2) / (yMAX - yMIN))
}
func xPointFloat32(x float32) float32 {
	return (x - xMIN) * (width / 2) / (xMAX - xMIN)
}
func yPointFloat32(y float32) float32 {
	return (y - yMIN) * (height / 2) / (yMAX - yMIN)
}

func main() {
	svg()
	pngg()
}

func svg() {
	var x, y float32 //[]float32
	f, _ := os.Create("image.svg")
	defer f.Close()
	fmt.Fprintf(f, "<svg viewBox='-%d 0 %d %d' xmlns='http://www.w3.org/2000/svg'>\n", width/2, width, height)

	//upLeft := image.Point{0, 0}
	//lowRight := image.Point{width, height}
	//img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	//cyan := color.RGBA{100, 200, 200, 0xff}

	x = 0 //append(x, 0)
	y = 0 //append(y, 0)
	curr := 0

	for i := 1; i < width*height/10; i++ {
		z := rand.Intn(101)
		if z == 1 {
			x = 0 //append(x, 0)
			y = 0.16 * y
			fmt.Fprintf(f, "<circle cx='%g' cy='%g' r='%g' fill='green'/>\n", xPointFloat32(x), yPointFloat32(y), radius) //x*160, y*100, radius)
			//fmt.Fprintf(f, "<polygon points='%g,%g'/>\n", x, y)
			//img.SetRGBA(xPoint(x), yPoint(y), cyan) //  S et(x, y, cyan)
		}
		if z >= 2 && z <= 86 {
			x = 0.85*x + 0.04*y                                                                                           //append(x, 0.85*x[curr]+0.04*y[curr])
			y = -0.04*x + 0.85*y + 1.6                                                                                    //append(y, -0.04*x[curr]+0.85*y[curr]+1.6)
			fmt.Fprintf(f, "<circle cx='%g' cy='%g' r='%g' fill='green'/>\n", xPointFloat32(x), yPointFloat32(y), radius) //x*160, y*100, radius)
			//fmt.Fprintf(f, "<polygon points='%g,%g'/>\n", x, y)
			//img.SetRGBA(xPoint(x), yPoint(y), cyan)
		}
		if z >= 87 && z <= 93 {
			x = 0.2*x - 0.26*y                                                                                            //append(x, 0.2*x[curr]-0.26*y[curr])
			y = 0.23*x + 0.22*y + 1.6                                                                                     //append(y, 0.23*x[curr]+0.22*y[curr]+1.6)
			fmt.Fprintf(f, "<circle cx='%g' cy='%g' r='%g' fill='green'/>\n", xPointFloat32(x), yPointFloat32(y), radius) //x*160, y*100, radius)
			//fmt.Fprintf(f, "<polygon points='%g,%g'/>\n", x, y)
			//img.SetRGBA(xPoint(x), yPoint(y), cyan)
		}
		if z >= 94 && z <= 100 {
			x = -0.15*x + 0.28*y                                                                                          //append(x, -0.15*x[curr]+0.28*y[curr])
			y = 0.26*x + 0.24*y + 0.44                                                                                    //append(y, 0.26*x[curr]+0.24*y[curr]+0.44)
			fmt.Fprintf(f, "<circle cx='%g' cy='%g' r='%g' fill='green'/>\n", xPointFloat32(x), yPointFloat32(y), radius) //x*160, y*100, radius)
			//fmt.Fprintf(f, "<polygon points='%g,%g'/>\n", x, y)
			//img.SetRGBA(xPoint(x), yPoint(y), cyan)
		}
		curr = curr + 1
	}
	fmt.Fprint(f, "</svg>")
}

func pngg() {
	var x, y float32 //[]float32

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width / 2, height / 2}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	cyan := color.RGBA{100, 200, 200, 0xff}

	x = 0 //append(x, 0)
	y = 0 //append(y, 0)

	curr := 0

	for i := 1; i < width*height; i++ {
		z := rand.Intn(101)
		if z == 1 {
			x = 0                                         //append(x, 0)
			y = 0.16 * y                                  //append(y, 0.16*y[0])
			img.SetRGBA(xPointInt(x), yPointInt(y), cyan) //  S et(x, y, cyan)
		}
		if z >= 2 && z <= 86 {
			x = 0.85*x + 0.04*y        //append(x, 0.85*x[curr]+0.04*y[curr])
			y = -0.04*x + 0.85*y + 1.6 //append(y, -0.04*x[curr]+0.85*y[curr]+1.6)
			img.SetRGBA(xPointInt(x), yPointInt(y), cyan)
		}
		if z >= 87 && z <= 93 {
			x = 0.2*x - 0.26*y        //append(x, 0.2*x[curr]-0.26*y[curr])
			y = 0.23*x + 0.22*y + 1.6 //append(y, 0.23*x[curr]+0.22*y[curr]+1.6)
			img.SetRGBA(xPointInt(x), yPointInt(y), cyan)
		}
		if z >= 94 && z <= 100 {
			x = -0.15*x + 0.28*y       //append(x, -0.15*x[curr]+0.28*y[curr])
			y = 0.26*x + 0.24*y + 0.44 //append(y, 0.26*x[curr]+0.24*y[curr]+0.44)
			img.SetRGBA(xPointInt(x), yPointInt(y), cyan)
		}
		curr = curr + 1
	}
	f, _ := os.Create("image.png")
	png.Encode(f, img)
	//println(x, y)
}
