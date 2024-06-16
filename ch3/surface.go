package ch3

import (
	"bytes"
	"fmt"
	"math"
)

const (
	width, height = 1200, 640
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.1
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func Surface() string {
	var buf bytes.Buffer
	buf.WriteString(`<!DOCTYPE html><html lang='en'><head><meta charset='UTF-8'><meta name='viewport' content='width=device-width, initial-scale=1.0'><title>Document</title></head><body>`)
	buf.WriteString(fmt.Sprintf("<svg viewBox='0 0 640 1200' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d' >", width, height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(float64(i+1), float64(j))
			bx, by := corner(float64(i), float64(j))
			cx, cy := corner(float64(i), float64(j+1))
			dx, dy := corner(float64(i+1), float64(j+1))
			color := color(float64(i+1), float64(j+1))
			buf.WriteString(fmt.Sprintf("<polygon fill='%s' points='%g, %g, %g, %g, %g, %g, %g, %g'/>\n", color, ax, ay, bx, by, cx, cy, dx, dy))
		}
	}

	buf.WriteString(fmt.Sprintln("</svg></body></html>"))

	return buf.String()
}

const (
	max          = 0xff0000
	min          = 0x0000ff
	colorStep    = (max - min) / 200
	colorAverage = (max - min) / 2
)

func corner(i float64, j float64) (float64, float64) {
	x := xyrange * (i/cells - 0.5)
	y := xyrange * (j/cells - 0.5)
	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r)
}

func color(i float64, j float64) string {
	x := xyrange * (i/cells - 0.5)
	y := xyrange * (j/cells - 0.5)
	z := (int(f(x, y) * 100))
	color := colorAverage + (z * colorStep)
	return fmt.Sprintf("#%.6x", color)
}
