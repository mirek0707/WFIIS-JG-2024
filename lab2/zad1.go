package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var stringMap = make(map[int][]string)

	for i := 0; i < 100; i++ {
		var length = rand.Intn(11) + 3
		var word = make([]uint8, length)
		j := 0
		for j < length {
			word[j] = letters[rand.Intn(len(letters))]
			if j > 0 {
				if word[j] == word[j-1] {
					break
				}
			}
			j++
		}
		if j != length {
			// fmt.Println("incorrect:", string(word), length)
			continue
		}
		// fmt.Println(string(word), length)
		stringMap[length] = append(stringMap[length], string(word))

	}
	for k, v := range stringMap {
		fmt.Println(k, v)
	}
	fmt.Println()
	for k, v := range stringMap {
		fmt.Println(k, v)
	}

}
