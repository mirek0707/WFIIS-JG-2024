package main

import (
	"fmt"
)

func addPrefix(prefix string) func(string) string {
	return func(str string) string { return prefix + str }
}

func addSuffix(prefix string) func(string) string {
	return func(str string) string { return str + prefix }
}

func aggregateDecorators(dec ...func(string) string) func(string) string {
	return func(str string) string {
		res := str
		for i := range dec {
			res = dec[i](res)
		}
		return res
	}
}

func main() {
	prefixDecorator := addPrefix("aaa")
	suffixDecorator := addSuffix("ccc")
	suffix2Decorator := addSuffix("ddd")
	aggDecorators := aggregateDecorators(prefixDecorator, suffixDecorator, suffix2Decorator)

	fmt.Println(prefixDecorator("bbb"))
	fmt.Println(suffixDecorator("bbb"))
	fmt.Println(aggDecorators("bbb"))
}
