package main

import "fmt"

func main() {
	str := "Hello 월드"    // 한영 섞인 문자열
	runes := []rune(str) // []rune으로 타입변환

	fmt.Printf("len(str) = %d\n", len(str))     // string 타입 길이
	fmt.Printf("len(runes) = %d\n", len(runes)) // []rune 타입 길이
}
