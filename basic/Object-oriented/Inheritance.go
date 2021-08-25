package main

import "fmt"

type PersoN struct {
	name string
	age int
	sex string
}

type Student struct {
	PersoN
	schoolName string
}

func main() {
	p1 := PersoN{"Fe", 20, "male"}
	fmt.Println(p1)
	fmt.Println("-----------------------")

	s1 := Student{p1, "NCHU"}
	printInfo(s1)

	s2 := Student{PersoN{"Luder", 21, "male"}, "NCHU"}
	printInfo(s2)

	s3 := Student{
		PersoN: PersoN{
			name: "Fe",
			age: 20,
			sex: "male",
		},
		schoolName: "NCHU",
	}
	printInfo(s3)

	s4 := Student{}
	s4.name = "Fe"
	s4.sex = "male"
	s4.age = 20
	s4.schoolName = "NCHU"
	printInfo(s4)
}

func printInfo(s1 Student) {
	fmt.Println(s1)
	fmt.Printf("%+v\n", s1)
	fmt.Printf("Name: %s, Age: %d, Gender: %s, School: %s\n", s1.name, s1.age, s1.sex, s1.schoolName)
	fmt.Println("-----------------------")
}