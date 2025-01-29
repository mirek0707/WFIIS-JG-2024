package main

import (
	"fmt"
	// "gonum.org/v1/plot"
	// "gonum.org/v1/plot/plotter"
	// "gonum.org/v1/plot/vg"
	// "image/color"
)

// func drawGraph(m map[int][]int){
// 	plt := plot.New()
// 	plt.Title.Text = ""
// 	plt.X.Label.Text = ""
// 	plt.Y.Label.Text = ""

// 	pts := make(plotter.XYs, liczba_punktów)

// 	for i := 0; i < liczba_punktów; i++ {
// 			pts[i].X = ...
// 			pts[i].Y = ...
// 	}

// 	line, _ := plotter.NewLine(pts)
// 	line.LineStyle.Color = color.RGBA{R: 1, G: 1, B: 1, A: 255}
// 	plt.Add(line)
// 	plt.Save(8*vg.Inch, 4*vg.Inch, name)
// }

func main() {
	g1 := Graph{[]*Node{}}
	g1.initGraph(5)
	g1.printGraph()

	mapIn := g1.getGraphDegreeInMap()
	fmt.Println("Degree in of nodes:")
	fmt.Println(mapIn)

	mapOut := g1.getGraphDegreeOutMap()
	fmt.Println("Degree out of nodes:")
	fmt.Println(mapOut)

	fmt.Println("\nShortest path algorithm:")
	d := g1.Floyd_Warshall_alg()
	for i := range d {
		for j := range d[i] {
			fmt.Print(d[i][j])
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}
