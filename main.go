package main

import (
	"flag"
	"fmt"
	"study/donovan/ch4"
)

// "crypto"

// "fmt"
// "study/donovan/ch1"
// "study/donovan/ch2"
// tempconv0 "study/donovan/ch2/tempconv"

// "fmt"
// "strconv"
// "strings"
// "study/donovan/ch1"

var path = flag.String("path", " ", "path to file to be read")

func main() {
	// ch1.Dup1()
	// ch1.DupFile()
	// ch1.Lissajous(os.Stdout)
	// ch1.Fecth(os.Stdout)
	// ch1.FetchAll()
	// ch1.Run()
	// ch1.Run2()
	// str := []byte("AabZ")
	// fmt.Println(str)
	// arr := make([]int, 0)
	// for _, v := range str {
	// 	fmt.Println(v)
	// 	for _, i := range strings.Split(fmt.Sprintf("%b", v), "") {
	// 		if val, err := strconv.Atoi(i); err == nil {
	// 			arr = append(arr, val)
	// 		} else {
	// 			fmt.Println(err)
	// 			continue
	// 		}
	// 	}
	// }
	// fmt.Println(arr, len(arr))
	// var i = 1
	// i1 := ch2.FunPointer(&i)
	// fmt.Println(i, i1)
	// ch2.Flags()
	// f := tempconv0.FToC(212)
	// fmt.Println(f.String())
	// // побитовые операции
	// // fmt.Printf("%08b\n", 33|4)
	// // fmt.Printf("%08b\n", 5^4)
	// // fmt.Printf("%08b\n", 5&4)
	// // fmt.Printf("%08b\n", 5&^4)
	// fmt.Println('A', 'a')

	// fmt.Println(ch3.Anagram("pots", "stop"))
	// sh1 := sha256.Sum256([]byte("x"))
	// sh2 := sha256.Sum256([]byte("X"))
	// fmt.Printf("%x,\n%x,\n%8.b,\n %8.b\n %t\n", sh1, sh2, sh1, sh2, sh1 == sh2)
	// fmt.Println(ch4.PC32(sh1, sh2))
	// fmt.Println(popcount.BitsDiff(&sh1, &sh2))
	// ch4.Shaalg()

	//	i := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	//	ch4.Reverse(&i)
	//	fmt.Println(i)
	//	j := []int{0, 1, 2, 3, 4, 5, 6}
	//	ch4.Rotate(j, 4)
	//	s := []string{"a", "a", "b", "b", "c", "d"}
	//	ch4.RemoveNDuplicate(s)

	// ch4.Dedup()
	// ch4.Charcount(*path)
	// ch4.CharcountByType(*path)
	// ch4.WordFreq(*path)
	// ch4.GetPoster()
	// ch1.Fecth(os.Stdout)
	// ch5.FindLinkMain()
	// ch5.CountNodes()
	// ch5.FindLinks2Main()
	// ch5.FindLinkMain()
	// opened, err := os.Open("chlen.html")
	// if err != nil {
	// 	os.Exit(1)
	// }
	// if doc, err := html.Parse(opened); err == nil {
	// var startEl func(node *html.Node) bool
	// var endEl func(node *html.Node) bool
	// var depth int
	// startEl = func(node *html.Node) bool {
	// 	if node.Type == html.ElementNode {
	// 		depth++
	// 		if node.NextSibling == nil {
	// 			fmt.Printf("%*s<%s>\n", depth*2, " ", node.Data)

	// 		} else {
	// 			fmt.Printf("%*s<%s/>\n", depth*2, " ", node.Data)

	// 		}
	// 	}
	// 	return true
	// }

	// endEl = func(node *html.Node) bool {
	// 	if node.Type == html.ElementNode {
	// 		if node.FirstChild != nil {
	// 			fmt.Printf("%*s</%s>\n", depth*2, " ", node.Data)
	// 		}
	// 		depth--
	// 	}
	// 	return true
	// }
	// ch5.ForEachNode(doc, func(node *html.Node) bool {
	// 	if node = ch5.ElementById(node, "hui"); node == nil {
	// 		return true
	// 	} else {
	// 		fmt.Println("huuue", *node)
	// 		return false
	// 	}
	// }, nil)
	// fmt.Println(ch5.Toposort(ch5.Prereqs))
	// flag.Parse()
	// fmt.Println(ch5.Extract(*path))
	// ch5.BreadthFirst(ch5.Crawl, os.Args[1:])

	// for v, k := range dirs {

	// }

	// fmt.Println(ch5.MinLogN(12, 2, 3, 10, 1, 1000, 200, 10, 11, 15, 20, 31, 44, 55))
	// ch5.Title(os.Args[1])
	// ch5.TraceExample()
	t := ch4.Tree{Value: 3, Left: &ch4.Tree{Value: 1}, Right: &ch4.Tree{Value: 2}}
	fmt.Println(t)
}
