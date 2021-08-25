package main

import "fmt"

type Address struct {
	province, city string
}

type Person struct {
	name string
	age int
	address *Address
}

func main() {
	p := Person{}
	p.name = "Fe"
	p.age = 20
	addr := Address{}
	addr.province = "北京市"
	addr.city = "海淀區"
	p.address = &addr
	fmt.Println("Name:", p.name, "Age:", p.age, "Province:", p.address.province, "City:", p.address.city)
	fmt.Println("--------------------")
	p.address.city = "昌平區"
	fmt.Println("Name:", p.name, "Age:", p.age, "Province:", p.address.province, "City:", p.address.city)
	fmt.Println("--------------------")
	p.address.city = "大興區"
	fmt.Println("Name:", p.name, "Age:", p.age, "Province:", p.address.province, "City:", p.address.city)
	fmt.Println("--------------------")
	p.address = &Address{
		province: "陝西省",
		city: "西安市",
	}
	fmt.Println(p)
	fmt.Println("Name:", p.name, "Age:", p.age, "Province:", p.address.province, "City:", p.address.city)
}