package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.IsNaN(3.4))
	fmt.Println(math.Ceil(1.000001))
	fmt.Println(math.Floor(1.999999))
	fmt.Println(math.Trunc(1.999999))
	fmt.Println(math.Abs(-1.3))
	fmt.Println(math.Max(-1.3, 0))
	fmt.Println(math.Min(-1.3, 0))
	fmt.Println(math.Dim(-12, -19))	// return max(-12 - (-19), 0)
	fmt.Println(math.Dim(-12, 19))	//	return max(-12 - 19, 0)
	fmt.Println(math.Mod(9, 4))	// return 9 % 4
	fmt.Println(math.Sqrt(9))
	fmt.Println(math.Cbrt(8))
	fmt.Println(math.Hypot(3, 4)) // return Sqrt(3*3 + 4*4)
	fmt.Println(math.Pow(2, 8))	// return 2^8
	fmt.Println(math.Log(1))			// return Loge(1)
	fmt.Println(math.Log2(16))
	fmt.Println(math.Log10(1000))
}