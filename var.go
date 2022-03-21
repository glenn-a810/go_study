package main

import "fmt"

func main() {
	var a int // 값을 저장할 변수1
	var b int // 값을 저장할 변수2

	n, err := fmt.Scan(&a, &b) // 입력 두 개 받기. &는 변수의 메모리 주소를 입력으로 넘김. 반환값 n은 성공적으로 입력한 값 개수
	if err != nil {
		fmt.Println(n, err) // 에러가 발생하면 에러코드 출력
	} else {
		fmt.Println(n, a, b) // 정상 입력되면 입력값 출력
	}
}
