package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var short []string
	var medium []string
	var long []string

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
		var short_cap int = cap(short)
		switch length {
		case 3, 4, 5, 6:
			short = append(short, string(word))
			if short_cap != cap(short) {
				fmt.Printf("Capacity of short changes from %d to %d. Actual length is %d.\n", short_cap, cap(short), len(short))
			}
		case 7, 8, 9:
			medium = append(medium, string(word))
		default:
			long = append(long, string(word))
		}

		// switch {
		// case length < 7:
		// 	short = append(short, string(word))
		// 	if short_cap != cap(short) {
		// 		fmt.Printf("Capacity of short changes from %d to %d. Actual length is %d.\n", short_cap, cap(short), len(short))
		// 	}
		// case length > 6 && length < 10:
		// 	medium = append(medium, string(word))
		// default:
		// 	long = append(long, string(word))
		// }
	}

	fmt.Println("\nShort length:")
	for i := range short {
		fmt.Println("\t", short[i], len(short[i]))
	}

	fmt.Println("\nMedium length:")
	for i := range medium {
		fmt.Println("\t", medium[i], len(medium[i]))
	}

	fmt.Println("\nLong length:")
	for i := range long {
		fmt.Println("\t", long[i], len(long[i]))
	}

}
