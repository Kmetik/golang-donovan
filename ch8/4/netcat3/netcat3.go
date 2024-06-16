package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var port = flag.Int("p", 8001, "port to connect")

func main() {
	flag.Parse()
	adress := fmt.Sprintf("localhost:%d", *port)
	connection, err := net.Dial("tcp", adress)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, connection)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(connection, os.Stdin)
	closeErr := connection.Close()
	fmt.Println(closeErr, "error")
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
