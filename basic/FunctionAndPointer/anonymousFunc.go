package main

import (
	"fmt"
	"math"
)

func main() {
	func(data string) {
		fmt.Println("Hello", data)
	}("Fe\n")

	f := func(data string) {
		fmt.Println(data)
	}
	f("Happy learning Go~\n")

	arr := []float64{1, 9, 16, 25, 31}
	visit(arr, func(v float64) {
		v = math.Sqrt(v)
		fmt.Printf("%.2f\n", v)
	})

	visit(arr, func(v float64) {
		v = math.Pow(v, 2)
		fmt.Printf("%.0f\n", v)
	})

}

func visit(list []float64, f func(float64)) {
	for _, value := range list {
		f(value)
	}
}