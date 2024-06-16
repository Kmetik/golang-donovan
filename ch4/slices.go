package ch4

import (
	"fmt"
)

func Reverse(arr *[10]int) {
	for i, j := 0, len(*arr)-1; i < len(*arr)/2; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func Rotate(arr []int, position int) {
	if position > len(arr)-1 {
		panic("ratatui")
	}
	for i := position; i >= 0; i-- {
		arr = append(arr, arr[i])
	}

	fmt.Println(arr[position+1:])
}

func RemoveNDuplicate(strings []string) {
	j := 0
	sLen := len(strings)
	for i := 1; i < sLen; i++ {
		if strings[j] != strings[i] {
			j++
			strings[j] = strings[i]
		}
	}
	fmt.Println(strings[:j+1])
}

// func ClueWhToASCII(s string) {
// 	runes := []rune(s)

// 	m := 0
// 	for i := 0; i < len(runes) - 1; i++ {
// 		if unicode.IsSpace(runes[i]) {
// 			m = i
// 		}
// 	}
// }

// func ReverseString(arr *[]byte) {
// 	for i, j := 0, len(*arr)-1; i < len(*arr)/2; i, j = i+1, j-1 {
// 		arr[i], arr[j] = arr[j], arr[i]
// 	}
// }
