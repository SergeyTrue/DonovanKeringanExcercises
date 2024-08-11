package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 700 //canvas size
	//number of grid cells
	xyrange = 30                  //axis ranges
	xyscale = width / 2 / xyrange //pixels per x or y unit
	zscale  = height * 0.4        //pixels per z unit
	angle   = math.Pi / 4
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var maxZ = math.Inf(-1)
var minZ = math.Inf(1)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "image/svg+xml")
		surface(writer, request)
	})
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func surface(out io.Writer, request *http.Request) {

	query := request.URL.Query()
	cells, _ := strconv.Atoi(query.Get("cells"))
	fmt.Println("cells=", cells)

	minZ = math.Inf(1)
	maxZ = math.Inf(-1)

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: green; stroke-width: 0.1' "+"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j, cells)
			bx, by, _ := corner(i, j, cells)
			cx, cy, _ := corner(i, j+1, cells)
			dx, dy, _ := corner(i+1, j+1, cells)
			if isValid(ax, ay, bx, by, cx, cy, dx, dy) {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'   />\n",
					ax, ay, bx, by, cx, cy, dx, dy, colorForHeight(az))
			}
		}
	}

	fmt.Fprintf(out, "</svg>\n")

}

func isValid(values ...float64) bool {
	for _, v := range values {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}
func corner(i, j, cells int) (float64, float64, float64) {

	//Find point (x,y) at corner of cell (i, j)
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)
	//compute surface height
	z := f(x, y)
	if z > maxZ {
		maxZ = z
	}
	if z < minZ {
		minZ = z
	}
	//Project (x, y, z isometrically onto 2-D SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyrange
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := rand.New(rand.NewSource(1000))
	// Parameters for the first egg
	a, b := 3., 3.
	var result float64
	for centerX := xyrange / -2.; centerX <= xyrange/2.; centerX += 5 {
		for centerY := xyrange / -2.; centerY <= xyrange/2.; centerY += 5 {
			dx := (x - centerX) / a
			dy := (y - centerY) / b
			ri := math.Hypot(dx, dy)
			result += r.Float64() * 2 * (math.Exp(-ri*3) * math.Sin(ri*2))
		}
	}
	return result
}
func colorForHeight(z float64) string {
	normalized := (z - minZ) / (maxZ - minZ)
	red := int(255 * normalized)
	blue := int(255 * (1 - normalized))
	return fmt.Sprintf("#%02x%02x%02x", red, 0, blue)
}
