package main

import "fmt"

func main() {
	var a int = 10
	var msg string = "변수"

	a = 20
	msg = "값 변경"
	fmt.Println(msg, a)
}
