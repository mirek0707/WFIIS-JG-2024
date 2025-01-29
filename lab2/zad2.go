package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	var sl = make([][]int, 10)
	for i := range sl {
		var length = rand.Intn(3) + 5
		sl[i] = make([]int, length)
		for j := range sl[i] {
			sl[i][j] = rand.Intn(6)
		}
		fmt.Println(sl[i])
	}

	var set = make(map[int]bool)

	for j := 0; j < 6; j++ {
		set[j] = true
		for i := range sl {
			if !slices.Contains(sl[i], j) {
				set[j] = false
			}
		}
	}

	for j := 0; j < 6; j++ {
		r, ok := set[j]
		if ok {
			if r {
				fmt.Println(j, "jest w każdym wierszu")
			} else {
				fmt.Println(j, "nie jest w każdym wierszu")
			}
		}

	}

}
