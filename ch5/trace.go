package ch5

import (
	"log"
	"time"
)

func bigSlowOperation() (res int) {
	defer trace("bigSlowOperation")()
	for i := 0; i < 10000000; i++ {
		res = i << 1
	}
	time.Sleep(10 * time.Second)
	return
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("вермя входа в операцию %s", msg)
	return func() {
		log.Printf("выход из операции %s (%s)", msg, time.Since(start))
	}
}
func TraceExample() {
	bigSlowOperation()
}
