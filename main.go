package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	println("listening")

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		panic(err)
	}
}
