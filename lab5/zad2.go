package main

import (
	// "bufio"
	"flag"
	"fmt"

	// "io"
	"log"
	"math"
	"math/rand"
)

var min = flag.Int("left", -5.0, "Left value")
var max = flag.Int("right", 5.0, "Right value")
var printHist = flag.Bool("printHist", true, "Print hist")

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}
func main() {
	flag.Parse()
	var hist = make(map[int]int)
	var arr []int
	var sum = 0
	for i := 0; i < 100; i++ {
		r := rand.Intn(*max-*min) + *min
		hist[r] += 1
		sum += r
		arr = append(arr, r)
	}
	var avr = float64(sum) / float64(100)
	fmt.Println("Average:", avr)

	var sd = 0.0
	for i := range arr {
		sd += math.Pow(float64(arr[i])-avr, 2)
	}
	sd = math.Sqrt(sd / 100)
	fmt.Println("Sd:", sd)
	if *printHist {
		fmt.Println(hist)
	}

}
