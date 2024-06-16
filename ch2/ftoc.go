package ch2

import "fmt"

func GetBoilingFreezing() {
	const (
		freezingF = 32.0
		boilingF  = 212.0
	)

	fmt.Printf("%fF = %fC", freezingF, ftoc(freezingF))
	fmt.Printf("%fF = %fC", freezingF, ftoc(boilingF))

}

func ftoc(f float32) float32 {
	return (f - 32) * 5 / 9
}

func FunPointer(i *int) int {
	*i++
	return *i
}
