package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/env", HelloEnv)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Full Cycle</h1>"))
}

func HelloEnv(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello, I'm %s, I'm %s.", name, age)
}
