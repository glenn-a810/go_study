package main

import "fmt"

func main() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b // 4266663.334329 -> float32 7자리 제한에 걸려 4266663으로 출력
	var d float32 = c * 3 // 12799990.002987 -> float32 7자리 제한에 걸려 12799989로 출력

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
