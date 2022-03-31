package main

import "fmt"

func main() {
	str := "Hello 월드!"
	arr := []rune(str)

	for i := 0; i < len(str); i++ {
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])
	}
}
