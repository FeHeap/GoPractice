package main

import (
	"fmt"
	"math"
)

func main() {

	//anonymous function
	res := func(a, b float64) float64 {
		return math.Pow(a, b)
	} (2,3)
	fmt.Println(res)


	//Anonymous Struct
	addr := struct {
		province, city string
	} {"陝西省", "西安市"}
	fmt.Println(addr)

	cat := struct {
		name, color string
		age int8
	} {
		name: "絨毛",
		color: "黑白",
		age: 1,
	}
	fmt.Println(cat)
}
