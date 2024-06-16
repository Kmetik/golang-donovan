package main

import (
	"fmt"
	"sort"
	sortpal "study/donovan/ch7/sortPal"
	sorttracks "study/donovan/ch7/sortTracks"
)

var tracks = []*sorttracks.Track{
	{Title: "Go", Author: "Moby", Year: 1992},
	{Title: "Go", Author: "Deliah", Year: 2012},
	{Title: "Run 2 Go", Author: "Alicia Keys", Year: 2017},
}

var p = []string{"a", "b", "c", "d", "b", "a"}
var is = []int{1, 2, 3, 5, 4, 3, 2, 1}

func main() {
	fmt.Println(sortpal.IsPalindrome(sort.IntSlice(is)))
}
