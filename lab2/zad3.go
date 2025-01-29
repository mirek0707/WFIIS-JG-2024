package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func insert(t []int, val int) ([]int, error) {
	for i := range len(t) {
		if t[i] == 0 {
			t[i] = val
			return t, nil
		} else if t[i] < val {
			continue
		} else {
			if t[len(t)-1] != 0 {
				return t, errors.New("tablica przepełniona")
			}
			for j := len(t) - 1; j > i; j-- {
				t[j] = t[j-1]
			}
			t[i] = val
			return t, nil
		}
	}
	return t, nil
}

func main() {
	var tab = [5]int{}
	for i := range 6 {
		var val = rand.Intn(100) + 1
		r, e := insert(tab[:], val)
		tab = [5]int(r)
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println("Tablica po iteracji", i, tab)
	}

}
