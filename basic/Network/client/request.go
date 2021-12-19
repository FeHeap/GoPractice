package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	testHttpNewRequest()
}

func testHttpNewRequest() {
	// 1. create a client
	client := http.Client{}
	// 2. create a request, request method may be GET, or POST
	request, err := http.NewRequest("GET", "https://zh.moegirl.org.cn/Mainpage", nil)
	CheckErr(err)
	// 3. client send request
	cookName := &http.Cookie{Name: "username", Value: "Steven"}
	// add cookie
	request.AddCookie(cookName)
	response, err := client.Do(request)
	CheckErr(err)
	// set request header
	request.AddCookie(cookName)
	defer response.Body.Close()
	// check data in request header
	fmt.Printf("Header: %+v\n", request.Header)
	fmt.Printf("status code: %v\n", response.StatusCode)
	// 4. manipulate data
	if response.StatusCode == 200 {
		data, err := ioutil.ReadAll(response.Body)
		fmt.Println("request success")
		CheckErr(err)
		fmt.Println(string(data))
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