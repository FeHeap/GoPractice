package main

import (
	"fmt"
	"strings"
)

type processFunc func(int) bool

func main() {
	result := StringToLower("AbcdefGHijklMNOPqrstUVWxyz", processCase)
	fmt.Println(result)
	result = StringToLower_("AbcdefGHijklMNOPqrstUVWxyz", processCase)
	fmt.Println(result)
	fmt.Println()

	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("odd element: ", odd)
	even := filter(slice, isEven)
	fmt.Println("even element: ",even)
}

func processCase(str string) string {
	result := ""
	for i, value := range str {
		if i % 2 == 0{
			result += strings.ToUpper(string(value))
		} else {
			result += strings.ToLower(string(value))
		}
	}
	return result
}

func StringToLower(str string, f func(string2 string) string) string {
	fmt.Printf("%T\n", f)
	return f(str)
}

type caseFunc func(string2 string) string

func StringToLower_(str string, f caseFunc) string {
	fmt.Printf("%T\n", f)
	return f(str)
}

func isEven(integer int) bool {
	if integer % 2 == 0 {
		return true
	}
	return false
}

func isOdd(integer int) bool {
	if integer % 2 == 0 {
		return false
	}
	return true
}

func filter(slice []int, f processFunc) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}