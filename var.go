package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string = "Hello"
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str)) // 구조체로 변경
	addr1 := stringHeader.Data                                    // Data 필드값을 addr1 변수로 저장

	str += " World"            // 문자열 추가
	addr2 := stringHeader.Data // Data 필드값을 addr2 변수로 저장

	str += " Welcome!"         // 또 문자열 추가
	addr3 := stringHeader.Data // Data 필드값을 addr3 변수로 저장

	// 메모리 주소값 출력
	fmt.Println(str)
	fmt.Printf("addr1:\t%x\n", addr1)
	fmt.Printf("addr2:\t%x\n", addr2)
	fmt.Printf("addr3:\t%x\n", addr3)
}
