package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type database map[string]dollars

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (d database) list(w http.ResponseWriter, r *http.Request) {
	for k, v := range d {
		fmt.Fprintf(w, "%s:%s \n", k, v)
	}
}

func (d database) updatePrice(item string, price float32) error {

	return nil
}

func (d database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	newPrice := r.URL.Query().Get("price")
	if price, ok := d[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "нет товара %s", item)
	} else {
		if len(newPrice) > 0 {
			if res, ok := strconv.ParseFloat(newPrice, 32); ok == nil {
				d[item] = dollars(res)
			} else {
				http.Error(w, "price should be number", http.StatusBadRequest)
				return
			}
		}
		log.Print(len(newPrice))
		fmt.Fprintf(w, "%v", price)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 32}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe("localhost:8001", http.DefaultServeMux))

}
