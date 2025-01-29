package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var fileFlag = flag.Bool("file", true, "Read from file or from stdin")
var delDuplFlag = flag.Bool("delDupl", true, "Delete duplicates")
var filterFlag = flag.String("filter", "", "Filter by given word")

func Read(r io.Reader) {
	var readedstrings []string
	in := bufio.NewScanner(r)

	for in.Scan() {
		line := in.Text()
		if len(line) == 0 {
			break
		}
		log.Println(line)
		readedstrings = append(readedstrings, string(line))
	}
	if *delDuplFlag {
		var set = make(map[string]bool)

		for j := range readedstrings {
			set[readedstrings[j]] = true
		}
		fmt.Println()
		fmt.Println("Without duplicates:")
		fmt.Println(set)
	}
	if *filterFlag != "" {
		fmt.Println()
		fmt.Println("Filtered strings")
		for j := range readedstrings {
			if strings.Contains(readedstrings[j], *filterFlag) {
				fmt.Println(readedstrings[j])
			}
		}

	}
	fmt.Println()
	fmt.Println("All strings:")
	fmt.Println(readedstrings)
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}
func main() {
	flag.Parse()

	log.Println("fileFlag has value:", *fileFlag, ". Value was set by user:", isFlagPassed("file"))
	log.Println("delDuplFlag has value:", *delDuplFlag, ". Value was set by user:", isFlagPassed("delDupl"))
	log.Println("filterFlag has value:", *filterFlag, ". Value was set by user:", isFlagPassed("filter"))

	if *fileFlag {
		file, err := os.Open("file.txt")
		if err != nil {
			fmt.Println(err)
		}
		Read(file)
	} else {
		Read(os.Stdin)
	}

}
