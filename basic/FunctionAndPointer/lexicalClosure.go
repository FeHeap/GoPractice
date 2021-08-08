package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\t", i)
		fmt.Println(add2(i, &sum))
	}
	fmt.Println()

	pos := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\t", i)
		fmt.Println(pos(i))
	}
	fmt.Println("----------------------")
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\t", i)
		fmt.Println(pos(i))
	}
	fmt.Println()

	myfunc := Counter()
	fmt.Println("myfunc", myfunc)
	fmt.Println(myfunc())
	fmt.Println(myfunc())
	fmt.Println(myfunc())
	myfunc_ := Counter()
	fmt.Println("myfunc_", myfunc_)
	fmt.Println(myfunc_())
	fmt.Println(myfunc_())

}

func add2(x int, sum *int) int {
	*sum += x
	return *sum
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		fmt.Printf("sum1 = %d\t", sum)
		sum += x
		fmt.Printf("sum2 = %d\t", sum)
		return sum
	}
}

func Counter() func() int {
	i := 0
	res := func() int {
		i += 1
		return i
	}
	fmt.Println("func in Counter():", res)
	return res
}