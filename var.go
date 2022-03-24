package main

import "fmt"

const PI = 3.14              // 타입 없는 상수
const FloatPI float64 = 3.14 // float64 타입 상수

func main() {
	var a int = PI * 100
	var b int = FloatPI * 100 // 타입 에러

	fmt.Println(a)
	fmt.Println(b)
}
