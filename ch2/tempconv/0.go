package tempconv0

import "fmt"

type Celsius float64
type Fahrenheit float64

var a int
var b int

func init() {
	a = 5
	b = 6
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

var (
	InitializedA = &a
	InitializedB = &b
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string { return fmt.Sprintf("%g C", c) }
