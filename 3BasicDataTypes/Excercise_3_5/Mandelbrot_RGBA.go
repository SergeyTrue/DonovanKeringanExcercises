package main

import (
	"github.com/crazy3lf/colorconv"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 10000, 10000
	)
	f, err := os.Create("mandelbrot.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			//Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(f, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 100
	const contrast = 250
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			h := float64(n%6) * 360. / 64
			s := 1.0
			v := 1 - float64(n)/400
			r, g, b, err := colorconv.HSVToRGB(h, s, v)
			if err != nil {
				log.Fatal(err)
			}
			return color.RGBA{r, g, b, contrast}
		}
	}
	return color.Black
}
