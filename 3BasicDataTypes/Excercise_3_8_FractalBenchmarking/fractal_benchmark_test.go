package main

import (
	"image"
	"testing"
)

func BenchmarkMandelbrot64(b *testing.B) {
	centerX, centerY, scale := -0.140009, 0.887, 1850000.
	for i := 0; i < b.N; i++ {
		drawImageToBenchmark(mandelbrot64, centerX, centerY, scale)
	}
}
func BenchmarkMandelbrot128(b *testing.B) {
	centerX, centerY, scale := -0.140009, 0.887, 1850000.
	for i := 0; i < b.N; i++ {
		drawImageToBenchmark(mandelbrot128, centerX, centerY, scale)
	}
}
func BenchmarkMandelbrotBigFloat(b *testing.B) {
	centerX, centerY, scale := -0.140009, 0.887, 1850000.
	for i := 0; i < b.N; i++ {
		drawImageToBenchmark(mandelbrotBigFloat, centerX, centerY, scale)
	}
}
func drawImageToBenchmark(fractalFunc FractalFunc, centerX, centerY, scale float64) {
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
}
