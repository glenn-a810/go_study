package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		for j := 1; j < 10; j++ {
			if i != j {
				continue
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}
}
