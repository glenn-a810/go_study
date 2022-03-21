package main

import "fmt"

func main() {
	var a int // 값을 저장할 변수1
	var b int // 값을 저장할 변수2

	//n, err := fmt.Scan(&a, &b) // 입력 두 개 받기. &는 변수의 메모리 주소를 입력으로 넘김. 반환값 n은 성공적으로 입력한 값 개수
	//n, err := fmt.Scanf("%d %d\n", &a, &b) // %d를 %s로 바꿔도 bad verb로 안됨. 변수 선언 자체가 int로 되어있음
	n, err := fmt.Scanln(&a, &b) // Scan()과 차이는 마지막 입력값 이후 반드시 enter키로 입력 종료를 해야함
	if err != nil {
		fmt.Println(n, err) // 에러가 발생하면 에러코드 출력
	} else {
		fmt.Println(n, a, b) // 정상 입력되면 입력값 출력
	}
}
