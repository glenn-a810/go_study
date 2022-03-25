package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	if false && IncreaseAndReturn() < 5 { // && 좌변이 false면 우변처리않고 false처리
		fmt.Println("1증가")
	}
	if true || IncreaseAndReturn() < 5 { // || 좌변이 true면 우변처리않고 true처리
		fmt.Println("2증가")
	}
	fmt.Println("cnt:", cnt)
}
