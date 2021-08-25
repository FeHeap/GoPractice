package main

import "fmt"

type User struct {
	string
	byte
	int8
	float64
}

func main() {
	user := User{"Fe", 'm', 20, 170.9}
	fmt.Println(user)
	fmt.Printf("Name: %s\n", user.string)
	fmt.Printf("Height: %.2f\n", user.float64)
	fmt.Printf("Gender: %c\n", user.byte)
	fmt.Printf("Age: %d\n", user.int8)
}
