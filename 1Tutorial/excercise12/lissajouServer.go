package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		lissajous(writer, request)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func lissajous(out io.Writer, request *http.Request) {
	query := request.URL.Query()
	cycles, _ := strconv.ParseFloat(query.Get("cycles"), 64)
	delay, _ := strconv.Atoi(query.Get("delay"))
	freq, _ := strconv.ParseFloat(query.Get("freq"), 64)
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
	const (
		res     = 0.001
		size    = 200
		nframes = 120
	)
	//freq := rand.Float64() * 4.0

	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		fmt.Println("Error while encoding: ", err)
		return

	}

}
