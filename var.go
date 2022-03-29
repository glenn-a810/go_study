package main

import "fmt"

type User struct {
	Name  string
	ID    string
	Age   int
	Level int // User의 Level 필드
}

type VIPUser struct {
	User
	Price int
	Level int // VIPUser의 Level 필드
}

func main() {
	user := User{"오드", "odd", 11, 10}
	vip := VIPUser{
		User{"루이", "luwie", 12, 20},
		250,
		3,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d 유저레벨: %d\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.Level,      // VIPUser의 Level
		vip.User.Level, // 포함된 구조체명을 쓰고 접근
	)
}
