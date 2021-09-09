package main

import (
	"fmt"
	"os"
)

func main() {
	fileName6 := "./test1"

	err := os.Remove(fileName6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s is removed!", fileName6)
	}

	err = os.RemoveAll(fileName6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s is removed!", fileName6)
	}
}
