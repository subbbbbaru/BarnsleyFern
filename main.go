package main

import (
	"errors"
	"flag"
	"log"

	"github.com/subbbbbaru/BarnsleyFern/fern"
)

var (
	flagFern         *int
	flagWidth        *int
	flagHeight       *int
	flagScale        *float64
	flagPoints       *int
	flagRenderFormat *int
)

const (
	Barnsley = iota
	Cyclosorus
	Modified
	Culcita
	Fishbone
	Tree
	Bee
	Your
)

func init() {
	flagRenderFormat = flag.Int("f", 1, "Render format")
	flagFern = flag.Int("fern", 0, `Format of fern:
	0: Barnsley
	1: Cyclosorus
	2: Modified
	3: Culcita
	4: Fishbone
	5: Tree
	6: Bee
	7: Your`)
	flagScale = flag.Float64("s", 0.85, "Format of fern scale")
	flagWidth = flag.Int("w", fern.PlotWidth, "Format of fern image width")
	flagHeight = flag.Int("h", fern.PlotHeight, "Format of fern image height")
	flagPoints = flag.Int("p", fern.Points, "Format of fern points")
}

func selectFern() (string, *[][]float64, error) {
	switch *flagFern {
	case Barnsley:
		return "Barnsley", &fern.Barnsley, nil
	case Cyclosorus:
		return "Cyclosorus", &fern.Cyclosorus, nil
	case Modified:
		return "Modified", &fern.Modified, nil
	case Culcita:
		return "Culcita", &fern.Culcita, nil
	case Fishbone:
		return "Fishbone", &fern.Fishbone, nil
	case Tree:
		return "Tree", &fern.Tree, nil
	case Bee:
		return "Bee", &fern.Bee, nil
	case Your:
		return "Your", nil, errors.New("soon")
	}
	return "", nil, errors.New("wrong flag")
}

func selectScale() {
	if *flagScale < 0 {
		log.Panic("Scale is not positive")
	}
	if *flagPoints < 0 {
		log.Panic("Points is not positive")
	}
}

func selectSize() {
	if *flagWidth < 0 {
		log.Panic("Width is not positive")
	}
	if *flagHeight < 0 {
		log.Panic("Height is not positive")
	}
}

func main() {
	flag.Parse()
	mmF := &fern.Ferns{}
	strFern, selectedFern, err := selectFern()
	if err != nil {
		log.Fatal(err.Error())
	}
	mmF.Fern = *selectedFern

	selectScale()
	selectSize()

	mmF.PlotHeight = *flagHeight
	mmF.PlotWidth = *flagWidth
	mmF.Scale = *flagScale
	mmF.PointS = *flagPoints

	switch *flagRenderFormat {
	case 1:
		mmF.GenPng(strFern)
	case 2:
		mmF.GenSvg(strFern)
	default:
		log.Panic("Error only 2 format(1: png, 2: svg)")
	}
}
