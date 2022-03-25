package main

import "fmt"

func getMyAge() (int, bool) {
	return 33, true
}

func main() {
	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("어림", age)
	} else if ok && age < 30 {
		fmt.Println("늙는중", age)
	} else if ok {
		fmt.Println("틀", age)
	} else {
		fmt.Println("Error")
	}

	fmt.Println("나이는 : ", age) // if 초기문에서 선언한 변수는 if문 안에서만 사용가능
}
