package main

import (
	"fmt"
	"strconv"
)

func f1[V ~[]T, T any](slice V, fun func(T) bool) []T {
	var result []T

	for _, item := range slice {
		if fun(item) {
			result = append(result, item)
		}
	}

	return result
}
func largerThanZero(n int) bool {
	return n > 0
}

func f2[V ~[]T, T any](slice V, fun func(T) T) []T {
	var result []T

	for _, item := range slice {
		result = append(result, fun(item))
	}
	return result
}

func f3[V ~[]T, T any](slice V, fun func(T, T) T) T {
	var result T

	for i, item := range slice {
		if i == 0 {
			result = item
		} else {
			result = fun(item, result)
		}
	}
	return result
}

func f4[T comparable, V any](m map[T]V) ([]T, []V) {
	var kResult []T
	var vResult []V
	for k, v := range m {
		kResult = append(kResult, k)
		vResult = append(vResult, v)
	}
	return kResult, vResult
}

type Pair[T comparable, V any] struct {
	key   T
	value V
}

func f5[T comparable, V any](m map[T]V) []Pair[T, V] {
	var result []Pair[T, V]
	for k, v := range m {
		result = append(result, Pair[T, V]{k, v})
	}
	return result
}

func f6[T comparable, V any](s1 []T, s2 []V) map[T]V {
	var s int
	m := make(map[T]V)
	if len(s1) > len(s2) {
		s = len(s2)
	} else {
		s = len(s1)
	}
	for i := 0; i < s; i++ {
		m[s1[i]] = s2[i]
	}
	return m
}

func f7[T int | float64](s []string) []T {
	var res []T
	var t interface{}
	t = *new(T)
	for i := range s {

		switch t.(type) {
		case int:
			in, err := strconv.Atoi(s[i])
			if err != nil {
				continue
			}
			res = append(res, T(in))
		case float64:
			in, err := strconv.ParseFloat(s[i], 64)
			if err != nil {
				continue
			}
			res = append(res, T(in))
		}
	}
	return res
}

func main() {
	numbers := []int{-1, -2, -3, 0, 5, 6, 7, 8, 9, -10}
	// 1
	ret := f1(numbers, largerThanZero)
	fmt.Println(ret)
	// 2
	ret = f2(numbers, func(n int) int { return n + 5 })
	fmt.Println(ret)
	// 3
	ret2 := f3(numbers, func(n1 int, n2 int) int { return n1 + n2 })
	fmt.Println(ret2)

	menu := make(map[string]float64)
	menu["eggs"] = 1.75
	menu["bacon"] = 3.22
	menu["sausage"] = 1.89
	// 4
	keys, values := f4(menu)
	fmt.Println(keys)
	fmt.Println(values)
	// 5
	res := f5(menu)
	fmt.Println(res)

	s1 := []string{"eggs", "bacon", "sausage"}
	s2 := []float64{1.75, 3.22, 1.89, 3.22}
	//6
	res2 := f6(s1, s2)
	fmt.Println(res2)

	//7
	num := []string{"-1", "-2", "-3", "0.6", "5.4"}
	res3 := f7[int](num)
	fmt.Println(res3)

}
