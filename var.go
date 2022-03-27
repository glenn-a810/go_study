package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for true { // 무한루프
		fmt.Println("입력하세요")
		var number int
		_, err := fmt.Scanln(&number) // 한 줄 입력받음
		if err != nil {               // 숫자가 아닌 경우
			fmt.Println("숫자를 입력하세요")

			// 키보드 버퍼를 삭제
			stdin.ReadString('\n')
			continue // 무한루프로 돌아감
		}
		fmt.Printf("입력하신 숫자는 %d입니다.\n", number)
		if number%2 == 0 { // 짝수 검사
			break // for문을 종료
		}
	}
	fmt.Println("for문이 종료되었습니다.")
}
