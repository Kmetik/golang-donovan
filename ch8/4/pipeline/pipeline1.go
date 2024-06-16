package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
		close(naturals)
	}()

	go func() {

		for x := range naturals {
			fmt.Println(naturals)
			squares <- x * x
		}
		close(squares)
	}()

	for i := range squares {
		fmt.Println(i)
	}
}
