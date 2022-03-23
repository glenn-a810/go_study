package main

import "fmt"

func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, "님 평균 점수는", avg, "입니다")
}

func main() {
	//math := 80
	//eng := 74
	//history := 95
	//fmt.Println("루이의 평균 점수는", (math+eng+history)/3, "입니다")
	//
	//math = 88
	//eng = 92
	//history = 53
	//fmt.Println("오드의 평균 점수는", (math+eng+history)/3, "입니다")
	//
	//math = 78
	//eng = 73
	//history = 78
	//fmt.Println("하루의 평균 점수는", (math+eng+history)/3, "입니다")

	PrintAvgScore("루이", 80, 74, 95)
	PrintAvgScore("오드", 88, 92, 53)
	PrintAvgScore("하루", 78, 73, 78)
}
