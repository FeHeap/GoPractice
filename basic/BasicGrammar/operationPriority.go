package main

import "fmt"

func main()  {
	var a int = 20
	var b int = 10
	var c int = 15
	var d int = 5
	var e int
	e = (a + b) * c / d		//(20 + 10) * 15 / 5
	fmt.Printf("(a + b) * c / d = %d\n", e)
	e = ((a + b) * c) / d	//((20 + 10) * 15) / 5
	fmt.Printf("((a + b) * c) / d = %d\n", e)
	e = (a + b) * (c / d)	//(20 + 10) * (15 / 5)
	fmt.Printf("(a + b) * (c / d) = %d\n", e)
	e = a + (b * c) / d		//20 + (10 * 15) / 5
	fmt.Printf("a + (b * c) / d = %d\n", e)
}
