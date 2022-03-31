package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHeader := (*reflect.StringHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str:\t%x\n", stringHeader.Data)
	fmt.Printf("slice:\t%x\n", sliceHeader.Data)
}
