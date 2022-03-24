package main

import "fmt"

func main() {
	const C int = 10 // 상수 선언

	var b int = C * 20
	C = 20          // 에러 - 상수는 대입문 좌변에 올 수 없음
	fmt.Println(&C) // 상수 C 주소출력
}
