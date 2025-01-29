package main

import (
	"fmt"
	"time"
)

func fibonacci(n uint64, a uint64, b uint64) uint64 {
	if n != 0 {
		return fibonacci(n-1, b, a+b)
	}
	return a
}

func print(delay time.Duration) {
	spinner := `-\|/`
	for {
		for _, char := range spinner {
			fmt.Printf("%c\r", char)
			time.Sleep(delay * time.Millisecond)
		}
	}

}

func main() {
	go print(50)
	fmt.Println(fibonacci(12000000, 0, 1))
}
