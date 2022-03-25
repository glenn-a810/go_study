package main

import "fmt"

//var price int

func HasRichFriend() bool {
	return true
}

func FriendCount() int {
	return 3
}

func main() {
	price := 35000

	if price > 50000 {
		if HasRichFriend() {
			fmt.Println("신발끈을 묶는다")
		} else {
			fmt.Println("돈을 나눠낸다")
		}
	} else if price >= 30000 && price <= 50000 {
		if FriendCount() > 3 {
			fmt.Println("신발끈을 묶는다")
		} else if FriendCount() <= 3 {
			fmt.Println("돈을 나눠낸다")
		}
	} else if price < 30000 {
		fmt.Println("내가 낸다")
	}
}
