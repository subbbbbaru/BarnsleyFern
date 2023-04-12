package main

import (
	"errors"
	"flag"
	"log"
	"subbbbbaru/fern"
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
	flagFern = flag.Int("fern", 1, "Format of fern")
	flagScale = flag.Float64("s", 0.85, "Format of fern scale")
	flagWidth = flag.Int("w", fern.PlotWidth, "Format of fern image width")
	flagHeight = flag.Int("h", fern.PlotHeight, "Format of fern image height")
	flagPoints = flag.Int("p", fern.Points, "Format of fern points")
}

func selectFern() (*[][]float64, error) {
	switch *flagFern {
	case Barnsley:
		return &fern.Barnsley, nil
		//myFern.Fern = fern.Barnsley
	case Cyclosorus:
		return &fern.Cyclosorus, nil
		//myFern.Fern = fern.Cyclosorus
	case Modified:
		return &fern.Modified, nil
		//myFern.Fern = fern.Modified
	case Culcita:
		return &fern.Culcita, nil
		//myFern.Fern = fern.Culcita
	case Fishbone:
		return &fern.Fishbone, nil
		//myFern.Fern = fern.Fishbone
	case Tree:
		return &fern.Tree, nil
		//myFern.Fern = fern.Tree
	case Bee:
		return &fern.Bee, nil
		//myFern.Fern = fern.Bee
	case Your:
		str := "Soon..."
		return nil, errors.New(str)
	}
	return nil, errors.New("Wrong flag")
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
	selectedFern, err := selectFern()
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
		mmF.GenPng()
	case 2:
		mmF.GenSvg()
	default:
		log.Panic("Error only 2 format(1: png, 2: svg)")
	}

}
