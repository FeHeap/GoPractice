package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	testHttpPost()
}

func testHttpPost() {
	// build parameter
	data := url.Values{
		"theCityName": {"重慶"},
	}
	// converse parameter to body
	reader := strings.NewReader(data.Encode())
	// send post request MIME format
	// http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName?thecityname=%E5%A4%A9%E6%B4%A5
	response, err := http.Post("http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName", "application/x-www-form-urlencoded", reader)
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
