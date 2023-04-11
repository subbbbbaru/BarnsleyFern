package main

import (
	"subbbbbaru/fern"
)

func main() {
	myFern := &fern.Ferns{
		Fern:       fern.Culcita,
		Scale:      2,
		PlotHeight: fern.PlotHeight,
		PlotWidth:  fern.PlotWidth,
		PointS:     fern.Points,
	}
	// myFern.GenPngg()
	// myFern.GenSvg()
	myFern.GenPng()
}
