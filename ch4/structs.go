package ch4

import (
	"bytes"
	"fmt"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

type Tree struct {
	Value       int
	Left, Right *Tree
}

func (t *Tree) String() string {
	fmt.Println("vizov")
	var buf bytes.Buffer
	for _, v := range appendValuesT(nil, t) {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		if buf.Len() == len("{") {
			continue
		}
		buf.WriteString(strconv.Itoa(v))
	}
	return buf.String()
}

func SortInt(intArr []int) []int {
	var root *tree
	// root value /v0 -6
	// root value 0 - 6, left
	for _, v := range intArr {
		root = add(root, v)
	}
	return appendValues(intArr[:0], root)
}

func appendValuesT(values []int, t *Tree) []int {
	if t != nil {
		values = appendValuesT(values, t.Left)
		values = append(values, t.Value)
		values = appendValuesT(values, t.Right)
	}
	return values
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = &tree{value: value}
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}
