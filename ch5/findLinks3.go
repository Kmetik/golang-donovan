package ch5

import "fmt"

// lookup := []string{"."}
// ch5.BreadthFirst(func(item string) []string {
// 	fmt.Println(item)
// 	entries, err := os.ReadDir(item)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	dirs := make([]string, 0, len(entries))

// 	for _, entry := range entries {
// 		if info, err := entry.Info(); err == nil {
// 			if info.IsDir() {
// 				dirs = append(dirs, fmt.Sprintf("%s/%s", item, info.Name()))
// 			} else {
// 				fmt.Println(info.Name())
// 			}
// 		}

// 	}
// 	return dirs
// }, lookup)

func BreadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if _, ok := seen[item]; !ok {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}

	}
}

func Crawl(url string) []string {
	fmt.Printf("url:%s\n", url)
	list, err := Extract(url)
	if err != nil {
		fmt.Println("Err:", err)
	}

	return list
}
