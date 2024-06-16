package ch2

import "fmt"

const boilingF = 212.0

func PrintBoiling() {
	f := boilingF
	c := (f - 32) * 5 / 9
	fmt.Printf("Farh: %f, Celsius: %f", f, c)
}
