package main

import (
    "net/http"
    "log"
    "fmt"
)

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "max-age=30, public, must-revalidate")
    v := r.URL.Query().Get("q")
		fmt.Println("get")
    fmt.Fprintf(w, "%s\n", v)
}
