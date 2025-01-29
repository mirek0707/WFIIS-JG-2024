package main

import (
	"errors"

	"regexp"
	"sort"
	"strings"
)

// func randString(length int) string {
// 	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	var word = make([]uint8, length)
// 	j := 0
// 	for j < length {
// 		word[j] = letters[rand.Intn(len(letters))]
// 		j++
// 	}
// 	return string(word)
// }

func splitString(s string, args ...string) ([]string, error) {
	var result []string
	if len(args) == 0 {
		sep := "[\t\n \\s]+"
		result = regexp.MustCompile(sep).Split(s, -1)

		return result, nil
	} else if len(args) == 1 {
		result = strings.Split(s, args[0])

		return result, nil
	} else {
		return result, errors.New("błędna ilość parametrów")
	}

}

func splitStringWithSort(s string, args ...string) ([]string, error) {
	var result []string
	if len(args) == 0 {
		sep := "[\t\n \\s]+"
		result = regexp.MustCompile(sep).Split(s, -1)
		sort.Strings(result)
		return result, nil
	} else if len(args) == 1 {
		result = strings.Split(s, args[0])
		sort.Strings(result)
		return result, nil
	} else {
		return result, errors.New("błędna ilość parametrów")
	}

}

func countOfNotEmptyLines(s string) (int, error) {
	result, err := splitString(s, "\n")
	if err != nil {
		return 0, err
	}
	var count int
	for i := range result {
		if result[i] != "" {
			count += 1
		}
	}
	return count, nil
}

func numberOfWords(s string) (int, error) {
	result, err := splitString(s)
	if err != nil {
		return 0, err
	}
	return len(result), nil
}

func numberOfMarks(s string) (int, error) {
	result, err := splitString(s, "")
	if err != nil {
		return 0, err
	}
	var count int
	for i := range result {
		if result[i] != " " && result[i] != "\n" && result[i] != "\t" {
			count += 1
		}
	}
	return count, nil
}

func wordsCount(s string) (map[string]int, error) {
	var stringMap = make(map[string]int)
	result, err := splitString(s)
	if err != nil {
		return stringMap, err
	}

	for i := range result {
		stringMap[result[i]] += 1
	}
	return stringMap, nil
}

func main() {

}
