package main

import (
	"fmt"
	"time"
)

func main() {
	delay := 100 * time.Millisecond
	go spinner(delay)
	fibS := fibFaster(45)
	fmt.Println(fibS)
	fibN := fib(45)
	fmt.Println(fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(c int) int {
	if c < 2 {
		return c
	}
	return fib(c-1) + fib(c-2)
}

func fibFaster(n int) int {
	s := make([]int, n+1)
	for i := 0; i < cap(s); i++ {
		if i < 2 {
			s[i] = i
		} else {
			s[i] = s[i-1] + s[i-2]
		}
	}

	return s[n]
}

// # (Public) Returns F(n).
// def fibonacci(n):
// 	if n < 0:
// 		raise ValueError("Negative arguments not implemented")
// 	return _fib(n)[0]

// # (Private) Returns the tuple (F(n), F(n+1)).
// def _fib(n):
// 	if n == 0:
// 		return (0, 1)
// 	else:
// 		a, b = _fib(n // 2)
// 		c = a * (b * 2 - a)
// 		d = a * a + b * b
// 		if n % 2 == 0:
// 			return (c, d)
// 		else:
// 			return (d, c + d)

func fibK(n int) int {
	if n < 0 {
		panic("Value should be more than zero")
	}
	return fibKPr(n)[0]
}

func fibKPr(n int) []int {
	if n == 0 {
		return []int{0, 1}
	} else {
		ab := fibKPr(n / 2)
		c := ab[0] * (ab[1]*2 - ab[0])
		d := ab[0]*ab[1] + ab[1]*ab[1]
		if n%2 == 0 {
			return []int{c, d}
		} else {
			return []int{d, c + d}
		}
	}
}
