package main

import (
	"github.com/crazy3lf/colorconv"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "server" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			drawImages(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	f, err := os.Create("mandelbrot_resampled.png")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	drawImages(f)
}

func drawSingleImage(out io.Writer) {
	srcImg := drawSrcMandelbrot()
	outImg := resample(1000, 1000, srcImg)
	err := png.Encode(out, outImg)
	if err != nil {
		log.Fatal(err)
	}
	err2 := png.Encode(out, srcImg)
	if err2 != nil {
		log.Fatal(err2)
	}
}

func drawImages(out io.Writer) {
	srcImg := drawSrcMandelbrot()
	resampledImg := resample(1000, 1000, srcImg)
	combinedWidth := srcImg.Bounds().Dx() + resampledImg.Bounds().Dy()
	combinedHeight := resampledImg.Bounds().Dy()
	combinedImg := image.NewRGBA(image.Rect(0, 0, combinedWidth, combinedHeight))
	draw.Draw(combinedImg, srcImg.Bounds(), srcImg, image.Point{0, 0}, draw.Src)
	resampledBounds := resampledImg.Bounds().Add(image.Point{srcImg.Bounds().Dx(), 0})
	draw.Draw(combinedImg, resampledBounds, resampledImg, image.Point{0, 0}, draw.Over)
	err := png.Encode(out, combinedImg)
	if err != nil {
		log.Fatal(err)
	}
}

func drawSrcMandelbrot() image.Image {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1000, 1000
	)

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
	return img
}
func mandelbrot(z complex128) color.Color {
	var v complex128
	const iterations = 200
	const contrast = 150
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			hue := float64(n%64) / 64 * 360.
			saturation := 1.0
			value := 1 - float64(n)/200
			r, g, b, err := colorconv.HSVToRGB(hue, saturation, value)
			if err != nil {
				log.Fatal(err)
			}
			return color.RGBA{r, g, b, contrast}
		}
	}
	return color.Black
}
func average9(x, y int, img image.Image) color.Color {
	var r, g, b uint32
	filters := []struct {
		dx, dy int
		weight uint32
	}{
		{-1, -1, 1}, {0, -1, 2}, {1, -1, 1},
		{-1, 0, 2}, {0, 0, 4}, {1, 0, 2},
		{-1, 1, 1}, {0, 1, 2}, {1, 1, 1},
	}
	for _, f := range filters {
		sr, sg, sb, _ := img.At(x+f.dx+1, y+f.dy).RGBA()
		r += sr * f.weight
		g += sg * f.weight
		b += sb * f.weight
	}
	return color.RGBA{uint8(r / 16 >> 8), uint8(g / 16 >> 8), uint8(b / 16 >> 8), 255}

}
func resample(xTarget, yTarget int, srcImage image.Image) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, xTarget, yTarget))
	for x := 0; x < xTarget; x++ {
		for y := 0; y < yTarget; y++ {
			img.Set(x, y, average9(x, y, srcImage))
		}
	}
	return img
}
