package main

import "fmt"

func funcAdd(x int, y int) int {
	return x + y
}

func sliceFunc(x []int) []int {
	return x

}

func main() {
	fmt.Println("hello world")
	//da
	fmt.Println(funcAdd(1, 2))
	fmt.Println("Слайс:", sliceFunc([]int{1, 2, 1, 2}))
}
