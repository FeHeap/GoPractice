package main

import (
	"fmt"
	"net/http"
)

func main() {
	testHttpGet()
}

func testHttpGet() {
	// get the server's data
	response, err := http.Get("http://www.baidu.com")
	CheckErr(err)
	fmt.Printf("status code: %v\n", response.StatusCode)
	if response.StatusCode == 200 {
		// manipulate StatusCode
		defer response.Body.Close()
		fmt.Println("request success")
		CheckErr(err)
	} else {
		fmt.Println("request fail", response.Status)
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