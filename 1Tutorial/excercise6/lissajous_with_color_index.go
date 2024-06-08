package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.RGBA{R: 0, G: 0, B: 0, A: 255},
	color.RGBA{R: 255, G: 0, B: 0, A: 255},
	color.RGBA{R: 0, G: 255, B: 0, A: 255},
}

const (
	blackIndex = 0
	redIndex   = 1
	greenIndex = 2
)

func lissajous(filepath string) {
	f, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	const (
		cycles  = 10
		res     = 0.001
		size    = 200
		nframes = 120
		delay   = 5
	)
	freq := rand.Float64() * 4.0

	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), redIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err = gif.EncodeAll(f, &anim)
	if err != nil {
		fmt.Println("Error while encoding: ", err)
		return

	}

}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	fmt.Println("Current directory:", wd)
	lissajous("my_gif.gif")
}
