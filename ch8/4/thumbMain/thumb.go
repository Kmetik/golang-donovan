package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"study/donovan/ch8/4/thumbMain/thumbnail"
	"sync"
)

func makeThumbnails(filenames []string) {
	for _, filename := range filenames {
		if _, err := thumbnail.ImageFile(filename); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, filename := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(filename)
	}

	for range filenames {
		<-ch
	}
}

func makeThumbnails5(filenames []string) ([]string, error) {
	type item struct {
		path string
		err  error
	}
	var thumbfiles []string
	ch := make(chan item, len(filenames))

	for _, filename := range filenames {
		go func(f string) {
			path, err := thumbnail.ImageFile(f)
			ch <- item{path, err}
		}(filename)
	}
	for range filenames {
		if item := <-ch; item.err != nil {
			return nil, item.err
		} else {
			thumbfiles = append(thumbfiles, item.path)
		}
	}
	return thumbfiles, nil
}

func makeThumbnails6(filenames <-chan string) int {
	var wg sync.WaitGroup
	sizes := make(chan int64)
	for filename := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			file, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(file)
			sizes <- info.Size()
		}(filename)
	}
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64

	for size := range sizes {
		total += size
	}
	fmt.Println(total)
	return int(total)
}
func main() {
	filenames := make(chan string)
	// var files = []string{"1.webp", "2.png", "3.png", "4.JPEG"}

	// go func() {
	// 	for _, file := range files {
	// 		filenames <- file
	// 	}
	// 	close(filenames)
	// }()

	input := bufio.NewScanner(os.Stdin)
	go func() {
		for input.Scan() {
			text := input.Text()
			filenames <- text
		}
		close(filenames)
	}()

	makeThumbnails6(filenames)

}
