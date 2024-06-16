package tempconv1

import (
	"flag"
	"fmt"
)

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

func (c Celsius) String() string { return fmt.Sprintf("%g ℃", c) }

type celsiusFlag struct {
	Celsius
}

func (c *celsiusFlag) Set(value string) error {
	var unit string
	var val float64

	fmt.Sscanf(value, "%f%s", &val, &unit)

	switch unit {
	case "C", "℃":
		c.Celsius = Celsius(val)
		return nil
	case "F":
		c.Celsius = FToC(Fahrenheit(val))
	default:
		return fmt.Errorf("temp should be C or ℃")

	}
	return nil
}

func CelsiucFlag(name string, defaultValue Celsius, description string) *Celsius {
	f := celsiusFlag{defaultValue}
	flag.CommandLine.Var(&f, name, description)
	return &f.Celsius
}
