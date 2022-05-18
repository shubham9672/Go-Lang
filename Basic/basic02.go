package main

import (
	"fmt"
	"sort"
)

func init() {
	fmt.Println("Hi Welcome to Basic Go lang Program")
}
func swap(a, b *int) {
	var c int
	c = *a
	*a = *b
	*b = c
}
func swap1(a, b int) (int, int) {
	return b, a
}
func sum(variable ...int) int {
	var c int = 0
	for _, v := range variable {
		c += v
	}
	// for i := 0; i < len(variable); i++ {
	// 	c += variable[i]
	// }
	return c
}
func Jointer(i func(p, q string) string) {
	fmt.Println(i("Hello ", " Year"))

}
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func basic02() {
	var arg int = 10
	var arg1 int = 20
	fmt.Println("Argument1: ", arg, "Argument2: ", arg1)
	c, d := swap1(arg, arg1)
	fmt.Println("Argument1: ", c, "Argument2: ", d, " After Swap")
	swap(&arg, &arg1)
	fmt.Println("Argument1: ", arg, "Argument2: ", arg1, " After Swap using pointer")
	fmt.Println(sum(1, 10, 15, 16))
	value := func(p, q string) string {
		return p + q + " 2022"
	}
	Jointer(value)
	var array [3]int
	fmt.Println(array)
	array[2] = 10
	array[1] = 20
	fmt.Println(array)
	s := []int{345, 78, 123, 10, 76, 2, 567, 5}
	sort.Ints(s)
	fmt.Println(s)
	list := make([]string, 2)
	list[0] = "Hi"
	list[1] = "Go"
	// list[2] = "Lang" it will give error
	list = append(list, "Lang")
	fmt.Println(list)
	fmt.Println(list[1:2])
	list1 := make([]int, len(s))
	copy(list1, s)
	fmt.Println(list1)
	m := make(map[int]string)
	m[1] = "one"
	m[2] = "two"
	m[30] = "thirty"
	// fmt.Println(m)
	for k, v := range m {
		fmt.Printf("%d -> %s\n", k, v)
	}
	fmt.Println(fact(5))
}
