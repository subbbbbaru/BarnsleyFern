package main

import (
	"subbbbbaru/fern"
)

func main() {
	myFern := fern.Ferns{
		Fern:       fern.Barnsley,
		Scale:      0.85,
		PlotHeight: fern.PlotHeight,
		PlotWidth:  fern.PlotWidth,
		PointS:     fern.Points,
	}
	myFern.GenPngg()
	myFern.GenSvg()
}
