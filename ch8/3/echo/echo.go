package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		panic(err)
	}
	for {
		if connection, err := listener.Accept(); err == nil {
			go handleConn(connection)
		} else {
			log.Fatal(err)
		}
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup

	input := bufio.NewScanner(c)
	for input.Scan() {
		text := input.Text()
		wg.Add(1)
		go func(s string) {
			go echo(c, s, 1000*time.Millisecond)
			defer wg.Done()
		}(text)
	}
	wg.Wait()
	if con, ok := c.(*net.TCPConn); ok {
		con.CloseWrite()
	} else {
		c.Close()
	}
}
