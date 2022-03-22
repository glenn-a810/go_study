package main

import "fmt"

func main() {
	var x1 int8 = 34   // 8bit 정수, 00100010
	var x2 int16 = 34  // 16bit 정수, 00000000 00100010
	var x3 uint8 = 34  // 8bit 부호가 없는 정수, 00100010
	var x4 uint16 = 34 // 16bit 부호가 없는 정수, 00000000 00100010

	fmt.Printf("^%d = %5d,\t %08b\n", x1, ^x1, uint8(^x1))
	fmt.Printf("^%d = %5d,\t %016b\n", x2, ^x2, uint16(^x2))
	fmt.Printf("^%d = %5d,\t %08b\n", x3, ^x3, ^x3)
	fmt.Printf("^%d = %5d,\t %016b\n", x4, ^x4, ^x4)
}
