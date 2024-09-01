package main

import (
	"fmt"
	"github.com/crazy3lf/colorconv"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/big"
	"math/cmplx"
	"net/http"
	"os"
	"strconv"
)

type FractalFunc func(complex128) color.Color

const (
	contrast   = 255
	iterations = 250
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "server" {

		handler := func(w http.ResponseWriter, r *http.Request) {
			params, fn := getQueryParams(r)
			drawImage(w, fn, params["centerX"], params["centerY"], params["scale"])
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe(":8000", nil))
		return
	}
	for _, fn := range []struct {
		filename string
		fn       FractalFunc
	}{
		{"mandelbrot64.png", mandelbrot64},
		{"mandelbrot128.png", mandelbrot128},
		{"mandelbrotBigFloat.png", mandelbrotBigFloat},
		//{"mandelbrotBigRat.png", mandelbrotBigRat},
	} {
		f, err := os.Create(fn.filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		drawImage(f, fn.fn, -0.140009, 0.887, 15)
		fmt.Printf("Image %s created\n", fn.filename)
	}
}
func colorize(n uint8) color.RGBA {
	hue := float64(n%64) / 64 * 360
	saturation := 1.0
	value := 1 - float64(n)/iterations
	r, g, b, err := colorconv.HSVToRGB(hue, saturation, value)
	if err != nil {
		log.Fatal(err)
	}
	return color.RGBA{r, g, b, contrast}

}

func drawImage(out io.Writer, fractalFunc FractalFunc, centerX, centerY, scale float64) {
	const (
		width, height = 1024, 1024
	)
	magnificationFactor := 1.0 / scale
	xmin := centerX - 2.0*magnificationFactor
	xmax := centerX + 2.0*magnificationFactor
	ymin := centerY - 2.0*magnificationFactor
	ymax := centerY + 2.0*magnificationFactor
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, fractalFunc(z))
		}

	}
	err := png.Encode(out, img)
	if err != nil {
		log.Fatal(err)
	}

}

func mandelbrot128(c complex128) color.Color {

	var z complex128
	for n := uint8(0); n < 200; n++ {

		z = z*z + c
		if cmplx.Abs(z) > 2 {
			return colorize(n)
		}
	}
	return color.Black
}

func mandelbrot64(c complex128) color.Color {
	var z complex64
	cc := complex64(c)
	for n := uint8(0); n < 200; n++ {
		z = z*z + cc
		if cmplx.Abs(complex128(z)) > 2 {
			return colorize(n)
		}
	}
	return color.Black
}

func mandelbrotBigFloat(c complex128) color.Color {

	zr, zi := big.NewFloat(0), big.NewFloat(0)
	cr, ci := big.NewFloat(real(c)), big.NewFloat(imag(c))
	two := big.NewFloat(2)
	radiusSq := new(big.Float).Mul(two, two)

	for n := uint8(0); n < iterations; n++ {
		zr2 := new(big.Float).Mul(zr, zr)
		zi2 := new(big.Float).Mul(zi, zi)
		re := new(big.Float).Sub(zr2, zi2)
		doubleZrZi := new(big.Float).Mul(new(big.Float).Mul(zr, zi), big.NewFloat(2))
		zr = new(big.Float).Add(re, cr)
		zi = new(big.Float).Add(doubleZrZi, ci)
		reAbs := new(big.Float).Add(zr2, zi2)
		if reAbs.Cmp(radiusSq) > 0 {
			return colorize(n)
		}
	}
	return color.Black
}

func mandelbrotBigRat(c complex128) color.Color {

	zr, zi := new(big.Rat), new(big.Rat)
	cr, ci := new(big.Rat).SetFloat64(real(c)), new(big.Rat).SetFloat64(imag(c))
	two := new(big.Rat).SetInt64(2)
	radiusSq := new(big.Rat).Mul(two, two)

	for n := uint8(0); n < iterations; n++ {
		zr2 := new(big.Rat).Mul(zr, zr)
		zi2 := new(big.Rat).Mul(zi, zi)
		re := new(big.Rat).Sub(zr2, zi2)
		doubleZrZi := new(big.Rat).Mul(new(big.Rat).Mul(zr, zi), two)
		zr = new(big.Rat).Add(re, cr)
		zi = new(big.Rat).Add(doubleZrZi, ci)
		reAbs := new(big.Rat).Add(zr2, zi2)
		if reAbs.Cmp(radiusSq) > 0 {
			return colorize(n)
		}
	}
	return color.Black
}

func getQueryParams(r *http.Request) (map[string]float64, func(complex128) color.Color) {
	params := map[string]float64{
		"centerX": -0.15,
		"centerY": 1.05,
		"scale":   15,
	}
	funcs := map[string]func(complex128) color.Color{
		"64":  mandelbrot64,
		"128": mandelbrot128,
		"bf":  mandelbrotBigFloat,
		"br":  mandelbrotBigRat,
	}
	centerX, err := strconv.ParseFloat(r.URL.Query().Get("centerX"), 64)
	if err == nil {
		params["centerX"] = centerX
	}
	centerY, err := strconv.ParseFloat(r.URL.Query().Get("centerY"), 64)
	if err == nil {
		params["centerY"] = centerY
	}
	scale, err := strconv.ParseFloat(r.URL.Query().Get("scale"), 64)
	if err == nil {
		params["scale"] = scale
	}
	method := r.URL.Query().Get("fn")

	function, exists := funcs[method]
	if !exists {
		function = mandelbrot128
	}
	fmt.Printf("method, %s\n", method)
	fmt.Printf("function, %T\n", function)
	return params, function
}
