package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

var ny = flag.String("NewYork", "localhost:8001", "NewYorkTime")
var ln = flag.String("London", "localhost:8002", "NewYorkTime")
var ty = flag.String("Tokyo", "localhost:8003", "NewYorkTime")

type city string

func (c city) Write(bytes []byte) (int, error) {
	n := len(bytes)
	value := string(bytes)
	fmt.Println(c, value)
	clocks[c] = value
	return n, nil
}

type cityClocks map[city]string

var clocks = make(cityClocks)

func main() {
	cities := map[city]*string{"NewYork": ny, "London": ln, "Tokyo": ty}

	for city, adress := range cities {
		go performConnection(city, adress)
	}

	for {
		// fmt.Fprint(os.Stdout, clocks)
		time.Sleep(1 * time.Second)
	}
}

func performConnection(c city, adress *string) {
	connection, err := net.Dial("tcp", *adress)
	if err != nil {
		return
	}
	defer connection.Close()
	handleConnection(connection, c)

}

func handleConnection(conn net.Conn, c city) {

	// io.Copy(clocks[city], conn)
}
