package main

import "net/http"

func main() {

	r := http.NewServeMux()
	r.Handle("/", http.FileServer(http.Dir("./static")))

	http.ListenAndServe("localhost:8080", r)
}
