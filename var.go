package main

import "fmt"

func main() {
	a := 324.13455
	c := 3.14

	fmt.Printf("%08.2f\n", a) // 최소 너비 8, 소수점 이하 2자리, 빈 너비에 0채음
	fmt.Printf("%08.2g\n", a) // 최소 너비 8, 총 출력되는 숫자 2자리, 빈 너비에 0채움
	// 324.13455에서 %08.3g면 00000324, 324를 2자리로 표현하기 위해 지수 표현으로 전환. 3.2e+02는 총 7자리므로 앞에 0채음
	fmt.Printf("%8.5f\n", a) // 최소 너비 8, 총 출력되는 숫자 5자리
	fmt.Printf("%f\n", c)    // 소수점 이하 6자리(디폴트)
}
