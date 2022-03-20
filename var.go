package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5

	var c int = int(b) // 정수->실수 타입변환하면 소수점 이하 버림 : 3
	d := float64(a * c)

	var e int64 = 7
	f := int64(d) * e

	var g int = int(b * 3) // 변환 전 3.5 x 3이 먼저 이루어 진 뒤 변환 : 10
	var h int = int(b) * 3 // 변환 후 3 x 3으로 연산 : 9

	fmt.Println(g, h, f)
}
