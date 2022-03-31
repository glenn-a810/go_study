package main

import "fmt"

func main() {
	str1 := "Hello\tWorld\n" // ""로 묶으면 특수문자가 동작함

	str2 := `Go is "awesome"!\nGo is simple and\t'powerful'` //``로 묶으면 특수문자가 동작하지 않음

	fmt.Println(str1)
	fmt.Println(str2)
}
