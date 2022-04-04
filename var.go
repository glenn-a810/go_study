package main

import "fmt"

func main() {
	var slice []int

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	slice[1] = 10 // 패닉 발생
	fmt.Println(slice)
}

//var slice1 = []int{1,2,3}
//var slice2 = []int{1,5:2,10:3} // [1 0 0 0 0 2 0 0 0 0 3]
//
//var array = [...]int{1,2,3} // 배열 선언. 길이가 3인 고정 길이 배열
//var slice = []int{1,2,3} // 슬라이스 선언
//
//var slice = make([]int, 3) // 길이 3개짜리 int 슬라이스
