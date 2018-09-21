package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var gen *rand.Rand = rand.New(rand.NewSource(makeTimestamp()))

func handler(w http.ResponseWriter, r *http.Request) {
	imgNo := gen.Intn(300)

	fmt.Fprintf(w, "<html><body><img src=\"https://www.catgifpage.com/gifs/%d.gif\"></body></html>\n", imgNo)
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}