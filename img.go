package main

import (
	"fmt"
	"net/http"
)

type Image struct {
	img
}

func name(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	NewImage(d, 100, 40).WriteTo(w)
	fmt.Println(ss)
}

func main() {
	http.HandleFunc("/", pic)
	s := &http.Server{
		Addr:           ":8100",
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
