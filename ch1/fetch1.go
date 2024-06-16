package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func Fecth(out io.Writer) {
	for _, arg := range os.Args[1:] {
		// http.SetCookie(http.Response{Status: 200}, &http.Cookie{Name: "SESSID", })
		url := arg
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
			fmt.Println(url)
		}
		resp, err := http.Get(url)
		if err != nil {
			os.Exit(1)
		}
		c, err := io.Copy(out, resp.Body)

		resp.Body.Close()

		if err != nil {
			continue
		}
		fmt.Printf("Statis is %d\n", resp.StatusCode)
		fmt.Printf("response is %d", c)
	}
}

func FetchAll() {
	start := time.Now()
	ch := make(chan string)

	for _, arg := range os.Args[1:] {
		// http.SetCookie(http.Response{Status: 200}, &http.Cookie{Name: "SESSID", })
		url := arg
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
			fmt.Println(url)
		}
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	time := time.Since(start).Seconds()
	fmt.Println(time)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	c, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	time := time.Since(start).Seconds()
	ch <- fmt.Sprintf("time %2f bytes: %7d url: %s", time, c, url)

}
