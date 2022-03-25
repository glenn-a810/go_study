package main

import "fmt"

func main() {
	var age = 22

	if age >= 10 && age <= 15 {
		fmt.Println("어리구만")
	} else if age > 30 || age < 20 {
		fmt.Println("20대는 아니네")
	} else {
		fmt.Println("늘금")
	}
}
