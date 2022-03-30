package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg Data) { // 매개변수로 Data를 받음
	arg.value = 999     // arg 데이터를 변경
	arg.data[100] = 999 // arg 데이터를 변경
}

func main() {
	var data Data

	ChangeData(data) // 인수로 data를 넣음
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}
