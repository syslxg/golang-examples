package main

import (
    //"fmt"
    "net/http"
	"io"
	"io/ioutil"
	"os"
)

func handler_stream(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("1.txt")
	io.Copy(w,file)
	defer file.Close()
}

func handler_load(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("1.txt")
	w.Write(data)
}

func main() {
    http.HandleFunc("/stream/", handler_stream)
    http.HandleFunc("/load/", handler_load)
    http.ListenAndServe(":8080", nil)
}