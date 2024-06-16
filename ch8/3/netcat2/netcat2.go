package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var port = flag.Int("port", 8001, "default port to connect")

func main() {
	flag.Parse()
	adress := fmt.Sprintf("localhost:%d", *port)
	fmt.Println(adress)
	connection, err := net.Dial("tcp", adress)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	go mustCopy(connection, os.Stdout)
	mustCopy(os.Stdin, connection)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
