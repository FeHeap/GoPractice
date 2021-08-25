package main

import "fmt"

type Employee struct {
	name, currency string
	salary float64
}

func main() {
	emp := Employee{"Daniel", "$", 2000}
	emp.printSalary()
	printSalary(emp)
}

// method
func (e Employee) printSalary() {
	fmt.Printf("Name: %s, Salary: %s%.2f\n", e.name, e.currency, e.salary)
}

// function
func printSalary(e Employee) {
	fmt.Printf("Name: %s, Salary: %s%.2f\n", e.name, e.currency, e.salary)
}
