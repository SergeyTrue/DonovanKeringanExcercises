package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 700            //canvas size
	cells         = 100                 //number of grid cells
	xyrange       = 30                  //axis ranges
	xyscale       = width / 2 / xyrange //pixels per x or y unit
	zscale        = height * 0.4        //pixels per z unit
	angle         = math.Pi / 4
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	cwd, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to get cwd:", err)
	}
	fmt.Println("cwd:", cwd)
	fileName := "output.svg"
	//filePath := filepath.Join(cwd, fileName)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Failed to create a file: ", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: green; stroke-width: 0.1' "+"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if isValid(ax, ay, bx, by, cx, cy, dx, dy) {
				fmt.Fprintf(file, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='white'   />\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}

	fmt.Fprintf(file, "</svg>\n")

}

func isValid(values ...float64) bool {
	for _, v := range values {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}
func corner(i, j int) (float64, float64) {
	//Find point (x,y) at corner of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	//compute surface height
	z := f(x, y)

	//Project (x, y, z isometrically onto 2-D SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyrange
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(2*r) / r * .3
}
