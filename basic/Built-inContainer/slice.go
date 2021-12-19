package main

import (
	"fmt"
	"strconv"
)

func main(){
	var numbers = make([]int, 3, 5)
	fmt.Printf("%T\n", numbers)
	fmt.Printf("len = %d, cap = %d, slice = %v\n\n", len(numbers), cap(numbers), numbers)


	s := []int{ 1, 2, 3 }
	fmt.Printf("%T\n", s)
	arr := [5]int{ 1, 2, 3, 4, 5 }
	s = arr[:]
	fmt.Printf("%T\n", s)
	fmt.Println()


	numbers = []int{ 0, 1, 2, 3, 4, 5, 6, 7, 8 }
	printSlice(numbers)
	fmt.Println("numbers ==", numbers)
	fmt.Println("numbers[1:4] ==", numbers[1:4])
	fmt.Println("numbers[:3] ==", numbers[:3])
	fmt.Println("numbers[4:] ==", numbers[4:])
	numbers_s1 := numbers[:2]
	printSlice(numbers_s1)
	numbers_s2 := numbers[2:5]
	printSlice(numbers_s2)
	fmt.Println()


	sliceCap()


	a := [4]float64{ 64.7, 89.8, 21, 78 }
	b := []int{ 2, 3, 5 }
	fmt.Printf("variable a -- address: %p, type: %T, value: %v, length: %d\n", &a, a, a, len(a))
	fmt.Printf("variable b -- address: %p, type: %T, value: %v, length: %d\n", &b, b, b, len(b))
	c := a
	d := b
	fmt.Printf("variable c -- address: %p, type: %T, value: %v, length: %d\n", &c, c, c, len(c))
	fmt.Printf("variable d -- address: %p, type: %T, value: %v, length: %d\n", &d, d, d, len(d))
	a[1] = 200
	fmt.Println("a =", a, "c =", c)
	d[0] = 100
	fmt.Println("b =", b, "d =", d)
	fmt.Println()


	array := [3]int{ 1, 2, 3}
	nums1 := array[:]
	nums2 := array[:]
	fmt.Println("array =", array)
	nums1[0] = 100
	fmt.Println("array =", array)
	nums2[1] = 200
	fmt.Println("array =", array)
	fmt.Println()

	fmt.Println("1. ----------")
	nums := make([]int, 0, 20)
	printSlices("nums:", nums)
	nums = append(nums, 0)
	printSlices("nums:", nums)
	nums = append(nums, 1)
	printSlices("nums:", nums)
	nums = append(nums, 2, 3, 4, 5, 6, 7)
	printSlices("nums:", nums)
	fmt.Println("2. ----------")
	s1 := []int{ 100, 200, 300, 400, 500, 600, 700 }
	nums = append(nums, s1...)
	printSlices("nums:", nums)
	fmt.Println("3. ----------")
	nums = nums[1:]
	printSlices("nums:", nums)
	nums = nums[:len(nums)-1]
	printSlices("nums:", nums)
	e := int(len(nums) / 2)
	fmt.Println("Middle number", e)
	nums = append(nums[:e], nums[e+1:]...)
	printSlices("nums:", nums)
	fmt.Println("4. ==========")
	nums_ := make([]int, len(nums), (cap(nums) * 2))
	count := copy(nums_, nums)
	fmt.Println("The number of numbers copied:", count)
	printSlices("nums_:", nums_)
	nums[len(nums)-1] = 99
	nums_[0] = 100.
	printSlices("nums_:", nums_)
	printSlices("nums:", nums)
	fmt.Println()


	var sa []string
	printSliceMsg(sa)
	for i := 0; i < 15; i++ {
		sa = append(sa, strconv.Itoa(i))
		printSliceMsg(sa)
	}
	printSliceMsg(sa)
}

func printSlice(x []int) {
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(x), cap(x), x)
}

func sliceCap() {
	arr0 := [...]string{ "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	fmt.Println("cap(arr0) =", cap(arr0), arr0)
	s01 := arr0[2:8]
	fmt.Printf("%T\n", s01)
	fmt.Println("cap(s01) =", cap(s01), s01)
	s02 := arr0[4:7]
	fmt.Println("cap(s02) =", cap(s02), s02)
	s03 := s01[3:9]
	fmt.Println("s03 := s01[3:9], s03 =", s03)
	s04 := s02[4:7]
	fmt.Println("s04 := s02[4:7], s04 =", s04)
	s04[0] = "x"
	fmt.Print(arr0, s01, s02, s03, s04)
	fmt.Println("\n")
}

func printSlices(name string, x []int) {
	fmt.Print(name, "\t")
	fmt.Printf("addr: %p \t len = %d \t cap = %d \t slice = %v\n", x, len(x), cap(x), x)
}

func printSliceMsg(sa []string) {
	fmt.Printf("addr: %p \t len: %d \t cap: %d \t value: %v\n", sa, len(sa), cap(sa), sa)
}