package main

import (
	"fmt"
	"net/http"
)

func webHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Demo go-webapp, this could be a spring app, or anything else</h1>")
}

func listenAndServe(port string) {
	fmt.Printf("serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/", webHandler)
	go listenAndServe("8080")
	select {}
}
