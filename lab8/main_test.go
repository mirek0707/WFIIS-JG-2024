package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	testTime = time.Now()
	res := m.Run()
	os.Exit(res)
}

func Test_splitString(t *testing.T) {
	res, err := splitString("aaagaaagffff", "g", "f")
	if err != nil {
		fmt.Println(err)
	}
	if len(res) != 0 {
		t.Error("error in splitString, returns not empty result", res, testTime)
	}
}

func Test_splitString2(t *testing.T) {
	result, err := splitString("aaagaaagffff", "g")
	if err != nil {
		t.Error("error in splitString:", err, testTime)
	} else {
		for i := range result {
			fmt.Println(result[i])
		}
	}
}

func Test_splitString3(t *testing.T) {
	result, err := splitString("aaa aaa ffff")
	if err != nil {
		t.Error("error in splitString:", err, testTime)
	} else {
		for i := range result {
			fmt.Println(result[i])
		}
	}
}

func Test_splitStringWithSort(t *testing.T) {
	result, err := splitStringWithSort("bbb aaa ffff")
	if err != nil {
		t.Error("error in splitStringWithSort:", err, testTime)
	} else {
		if result[0] != "aaa" {
			t.Error("error in splitStringWithSort:", result, testTime)
		}
	}
}
func Test_countOfNotEmptyLinest(t *testing.T) {
	result, err := countOfNotEmptyLines("bbbaaaffff\n\nddddd\n")
	if err != nil {
		t.Error("error in countOfNotEmptyLines:", err, testTime)
	} else {
		if result != 2 {
			t.Error("error in countOfNotEmptyLines:", result, testTime)
		}
	}
}
func Test_numberOfWords(t *testing.T) {
	result, err := numberOfWords("bbb aaa ffff\n\nddd dd")
	if err != nil {
		t.Error("error in numberOfWords:", err, testTime)
	} else {
		if result != 5 {
			t.Error("error in numberOfWords:", result, testTime)
		}
	}
}

func Test_numberOfMarks(t *testing.T) {
	result, err := numberOfMarks("bbb aaa ffff\n\nddd dd")
	if err != nil {
		t.Error("error in countOfNotEmptyLines:", err, testTime)
	} else {
		if result != 15 {
			t.Error("error in countOfNotEmptyLines:", result, testTime)
		}
	}
}
func Test_wordsCount(t *testing.T) {
	result, err := wordsCount("bbb aaa ffff\n\nddd dd")
	if err != nil {
		t.Error("error in wordsCount:", err, testTime)
	} else {
		if len(result) == 0 {
			t.Error("error in wordsCount:", result, testTime)
		}
	}
}

var blackhole []string

func BenchmarkFuns(b *testing.B) {
	for _, file := range []string{"Latin-Lipsum_5.txt", "Latin-Lipsum_13.txt", "Latin-Lipsum_20.txt"} {
		b.Run(file, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f, _ := os.ReadFile(file)
				str := string(f)
				res, _ := splitStringWithSort(str)
				blackhole = res
			}
		})
	}
}
