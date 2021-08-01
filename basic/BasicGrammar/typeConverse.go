package main

import "fmt"

func main()  {
	chinese := 90
	english := 80.9
	avg := (chinese + int(english)) / 2
	avg2 := (float64(chinese) + english) / 2
	fmt.Printf("%T , %d \n", avg, avg)		//int , 85
	fmt.Printf("%T , %f \n\n", avg2, avg2)		//float64 , 85.450000

	a := 97
	x := 19968
	result := string(a)
	fmt.Println(result)		//a
	result = string(x)
	fmt.Println(result)		//一
}