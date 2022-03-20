package main

import "fmt"

var g int = 10 // 패키지 전역 변수

func main() {
	var m int = 20 // 지역 변수

	{
		var s int = 50 // 지역 변수
		fmt.Println(m, s, g)
	} // s 지역 변수 사라짐

	fmt.Println(m, s, g) // s가 범위 내 없으므로 에러
}
