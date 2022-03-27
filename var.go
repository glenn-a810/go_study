package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for true {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}
}
