package main
import (
	"fmt" 
)


func main() {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	fmt.Printf("%T %T\n", letters, letters[0])
}