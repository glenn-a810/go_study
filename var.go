package main

import "fmt"

func main() {
	a := 123
	b := 456
	c := 123456789

	fmt.Printf("%5d, %5d\n", a, b)    // 최소 너비보다 짧은 값 너비
	fmt.Printf("%05d, %05d\n", a, b)  // 최소 너비보다 짧은 값 빈곳에 0채우기
	fmt.Printf("%-5d, %-05d\n", a, b) // 최소 너비보다 짧은 값 왼쪽으로 shift

	fmt.Printf("%5d, %5d\n", c, c)    // 최소 너비보다 긴 값 너비
	fmt.Printf("%05d, %05d\n", c, c)  // 최소 너비보다 긴 값 0채우기
	fmt.Printf("%-5d, %-05d\n", c, c) // 최소 너비보다 긴 값 왼쪽 shift
	// 최소 너비보다 긴 값은 지정한 최소 너비를 무시하고 그냥 길게 출력함
}
