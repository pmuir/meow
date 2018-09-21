package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var gen *rand.Rand = rand.New(rand.NewSource(makeTimestamp()))

func handler(w http.ResponseWriter, r *http.Request) {
	imgNo := gen.Intn(300)

	var payload = map[string]interface{} {
		"text" : "Here's a cat for you",
		"attachments" : [] map[string]string {
			{
				"image_url": fmt.Sprintf("https://www.catgifpage.com/gifs/%d.gif", imgNo),
			},
		},
	}
	json, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println(err)
	} else {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(200)
		w.Write(json)
	}
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}