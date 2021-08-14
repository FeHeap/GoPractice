package main

import "fmt"

func main() {
	username := ""
	age := 0
	fmt.Scanln(&username, &age)
	fmt.Println("帳號信息為:", username, age)
}
