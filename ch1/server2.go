package ch1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"study/donovan/ch3"
	"sync"
)

//54

var mu sync.Mutex

var count int

func Run2() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", handlerCount)

	http.HandleFunc("/liss", handleLiss)
	http.HandleFunc("/surface", handleSurface)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s", r.URL)
}

func handlerCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "%d", count)
	mu.Unlock()
}

func handleLiss(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles, err := strconv.ParseFloat(r.Form.Get("cycles"), 32)
	if err != nil {
		cycles = 5
	}
	Lissajous(w, cycles)

}

func handleSurface(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprint(w, ch3.Surface())
}
