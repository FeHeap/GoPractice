package main

import "fmt"

func main() {
	var nums_1 = [5]int{ 1, 2, 3, 4, 5 }
	var nums_2 = [...]int{ 1, 2, 3, 4, 5, 6 }
	nums_2[4] = 27
	fmt.Printf("The length of nums_1 is %d.\n", len(nums_1))
	fmt.Printf("The length of nums_2 is %d.\n", len(nums_2))
	fmt.Println("nums_1:")
	for i := 0; i < len(nums_1); i++ {
		fmt.Print(nums_1[i], "\t")
	}
	fmt.Println()
	fmt.Println("nums_2:")
	for _, value := range nums_2 {
		fmt.Print(value, "\t")
	}
	fmt.Println("\n")


	var arr2D = [5][2]int{ {0, 0}, {1, 2}, {3, 4}, {5, 6}, {0, 4} }
	fmt.Println(len(arr2D))
	fmt.Println(len(arr2D[0]))
	for i := 0; i < len(arr2D); i++ {
		for j := 0; j < len(arr2D[0]); j++ {
			fmt.Printf("arr2D[%d][%d] = %d\n", i, j, arr2D[i][j])
		}
	}
	fmt.Println()


	a := [...]string{ "USA", "China", "India", "Germany", "France" }
	b := a  // A copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
