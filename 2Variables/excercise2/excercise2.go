package main

import (
	"flag"
	"fmt"
	tempconv "github.com/SergeyTrue/DonovanKeringanExcercises/2Variables/excercise1"
	"os"
)

var from = flag.String("from", "", "from convert")
var to = flag.String("to", "", "convert to")
var temp = flag.Float64("temp", 0, "temperature")

func main() {
	flag.Parse()
	if len(os.Args) > 1 {
		if *from == "c" && *to == "f" {
			fmt.Printf("Вот мы и законвертили цельсий в файренгейт: %.2f\n", tempconv.CToF(tempconv.Celsius(*temp)))
		} else {
			fmt.Println("No arguments have been provided")
		}
	}
}
