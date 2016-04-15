package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello server, %q", html.EscapeString(r.URL.Path))
}

func main() {
    log.Println("Iniciando servidor...")
    http.HandleFunc("/", Hello)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
