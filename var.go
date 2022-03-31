package main

import "fmt"

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	slice[2] = 'a'

	fmt.Println(str)
	fmt.Printf("%s\n", slice) // str과 slice가 다른 메모리 주소를 가리키고 있음
}
