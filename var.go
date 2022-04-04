package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}

	//slice2 := append(slice, 4)
	slice2 := append(slice, 4, 5, 6, 7, 8) // 여러 값 추가

	fmt.Println(slice)
	fmt.Println(slice2)
}
