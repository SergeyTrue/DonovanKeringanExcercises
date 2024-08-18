package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"os"
)

const (
	iterations             = 200
	tolerance              = 1e-6
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 5000, 5000
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "server" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			drawImage(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	f, err := os.Create("newtonFractal.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	drawImage(f)
}

func drawImage(out io.Writer) {
	img := drawSrcNewton()

	err := png.Encode(out, img)
	if err != nil {
		log.Fatal(err)
	}
}

func drawSrcNewton() image.Image {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			//Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	return img
}
func polynominal(z complex128) complex128 {
	return cmplx.Pow(z, 6) - 1
}
func derivative(z complex128) complex128 {
	return 6. * cmplx.Pow(z, 5)
}

var rootColors = []color.RGBA{
	{255, 0, 0, 255},   // Red
	{0, 255, 0, 255},   // Green
	{0, 0, 255, 255},   // Blue
	{255, 255, 0, 255}, // Yellow
	{0, 255, 255, 255}, // Cyan
	{255, 0, 255, 255}, // Magenta
}

func closestRoot(z complex128) int {
	var roots = []complex128{
		cmplx.Exp(2i * math.Pi / 6 * 0),
		cmplx.Exp(2i * math.Pi / 6 * 1),
		cmplx.Exp(2i * math.Pi / 6 * 2),
		cmplx.Exp(2i * math.Pi / 6 * 3),
		cmplx.Exp(2i * math.Pi / 6 * 4),
		cmplx.Exp(2i * math.Pi / 6 * 5),
	}
	minDist := cmplx.Abs(roots[0] - z)
	closest := 0
	for i, root := range roots {
		dist := cmplx.Abs(root - z)
		if dist < minDist {
			minDist = dist
			closest = i
		}
	}
	return closest
}

func interpolateColor(c color.RGBA, n uint8) color.Color {
	factor := float64(10*n) / float64(iterations)
	return color.RGBA{
		uint8(math.Round(float64(c.R) * factor)),
		uint8(math.Round(float64(c.G) * factor)),
		uint8(math.Round(float64(c.B) * factor)),
		255,
	}
}

func newton(z complex128) color.Color {

	for n := uint8(0); n < iterations; n++ {
		dz := polynominal(z) / derivative(z)
		z -= dz
		if cmplx.Abs(dz) < tolerance {
			root := closestRoot(z)
			return interpolateColor(rootColors[root], n)
		}
	}
	return color.Black
}
