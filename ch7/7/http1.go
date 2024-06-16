package main

import (
	"fmt"
	"log"
	"net/http"
)

type database map[string]dollars

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/list":
		for k, v := range db {
			fmt.Fprintf(w, "%s :%v \n", k, v)
		}
	case "/price":
		item := r.URL.Query().Get("item")
		if price, ok := db[item]; !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "нет товара %s", item)
		} else {
			fmt.Fprintf(w, "%v", price)
		}

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "NOT FOUND")
	}

}

func main() {
	db := database{"shoes": 50, "socks": 32}

	log.Fatal(http.ListenAndServe("localhost:3080", db))
}
