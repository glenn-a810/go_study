package main

import "fmt"

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		//return 0, false
		result = 0
		success = false
		return
	} else {
		//return a / b, true
		result = a / b
		success = true
		return
	}
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}
