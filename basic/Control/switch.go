package main

import (
	"fmt"
)

func main()  {
	grade := ""
	score := 78.5
	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		grade = "D"
	default:
		grade = "E"
	}
	fmt.Printf("Your Level is: %s\n", grade)
	fmt.Print("mean ")
	switch grade {
	case "A":
		fmt.Println("excellent")
	case "B":
		fmt.Println("good")
	case "C":
		fmt.Println("normal")
	case "D":
		fmt.Println("not good")
	default:
		fmt.Println("bad")
	}

	year := 2008
	month := 2
	days := 0
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
			days = 29
		} else {
			days = 28
		}
	default:
		days = -1
	}
	fmt.Printf("There are %d days in %d.%d\n", days, month, year)

	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("The type of x is: %T\n", i)
	case int:
		fmt.Println("The type of x is int")
	case float64:
		fmt.Println("The type of x is float64")
	case func(int) float64:
		fmt.Println("The type of x is func(int)")
	case bool, string:
		fmt.Println("The type of x is bool or string")
	default:
		fmt.Println("unknown")
	}
}
