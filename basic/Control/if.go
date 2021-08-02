package main

import "fmt"

func main()  {
	num := 20
	if num % 2 == 0 {
		fmt.Println(num, "even")
	} else {
		fmt.Println(num, "odd")
	}

	score := 88
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else if score < 60 {
		fmt.Println("E")
	}

	if Num := 10; Num % 2 == 0 {
		fmt.Println(Num, "even")
	} else {
		fmt.Println(Num, "odd")
	}

	if Score := 98; Score >= 60 {
		if Score >= 70 {
			if Score >= 80 {
				if Score >= 90 {
					fmt.Println("A")
				} else {
					fmt.Println("B")
				}
			} else {
				fmt.Println("C")
			}
		} else {
			fmt.Println("D")
		}
	} else {
		fmt.Println("E")
	}
}
