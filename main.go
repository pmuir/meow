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
	animal := ""
	if r.URL.Path == "/meow" {
		animal = "cat"
	} else if r.URL.Path == "/woof" {
		animal = "dog"
	} else if r.URL.Path == "/" {
		w.WriteHeader(200)
		w.Write([]byte("Meow!"))
		return
	} else {
		w.WriteHeader(500)
		w.Write([]byte("Unknown animal!"))
		return
	}
	imgNo := gen.Intn(300)

	var payload = map[string]interface{} {
		"text" : fmt.Sprintf("Here's a %s for you", animal),
		"attachments" : [] map[string]string {
			{
				"image_url": fmt.Sprintf("https://www.%sgifpage.com/gifs/%d.gif", animal, imgNo),
			},
		},
	}
	json, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println(err)
		return
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