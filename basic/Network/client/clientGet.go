package main

import (
	"fmt"
	"net/http"
)

func main() {
	testClientGet()
}

func testClientGet() {
	// create a client
	client := http.Client{}
	// request through client
	response, err := client.Get("https://zh.moegirl.org.cn/Mainpage")
	CheckErr(err)
	fmt.Printf("status code: %v\n", response.StatusCode)
	if response.StatusCode == 200 {
		fmt.Println("request success")
		defer response.Body.Close()
	}
}

// check error
func CheckErr(err error) {
	defer func() {
		if ins, ok := recover().(error); ok {
			fmt.Println("ERROR: ", ins.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}