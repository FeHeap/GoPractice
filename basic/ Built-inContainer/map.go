package main

import "fmt"

func main() {
	var country = map[string]string{
		"China": "Beijing",
		"Japan": "Tokyo",
		"India": "New Delhi",
		"France": "Paris",
		"Italy": "Rome",
	}
	fmt.Println(country)
	rating := map[string]float64{"c": 5, "Go": 4.5, "Python": 4.5, "C++": 3}
	fmt.Println(rating)
	countryMap := make(map[string]string)
	countryMap["China"] = "Beijing"
	countryMap["Japan"] = "Tokyo"
	countryMap["India"] = "New Delhi"
	countryMap["France"] = "Paris"
	countryMap["Italy"] = "Rome"
	for k, v := range countryMap {
		fmt.Println("Country", k, "Capital", v)
	}
	fmt.Println("----------")
	for _, v := range countryMap {
		fmt.Println( "Capital", v)
	}
	fmt.Println("----------")
	for k := range countryMap {
		fmt.Println("Country", k, "Capital", countryMap[k])
	}
	fmt.Println("----------")

	value, ok := countryMap["England"]
	fmt.Printf("%q\n", value)
	fmt.Printf("%T, %v\n", ok, ok)
	if ok {
		fmt.Println("Capital", value)
	} else {
		fmt.Println("Can't find the capital in countryMap!")
	}
	if val, Ok := countryMap["USA"]; Ok {
		fmt.Println("Capital", val)
	} else {
		fmt.Println("Can't find the capital in countryMap!")
	}
	fmt.Println()


	map_ := map[string]string{
		"element": "div",
		"width": "100px",
		"height": "200px",
		"border": "solid",
		"background": "none",
	}
	fmt.Println("Before delete,", map_)
	if _, Ok := map_["background"]; Ok {
		delete(map_, "background")
	}
	fmt.Println("After delete,", map_)
	map_ = make(map[string]string)
	fmt.Println("After clean,", map_)
	fmt.Println()

	personSalary := map[string]int{
		"Steven":	18000,
		"Daniel": 	5000,
		"Josh":		20000,
	}
	fmt.Println("Original salary:", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["Daniel"] = 8000
	fmt.Println("After changing newPersonSalary:", newPersonSalary)
	fmt.Println("personSalary was also Affected:", personSalary)
}
