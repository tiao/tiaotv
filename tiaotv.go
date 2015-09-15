package main

import (
  "log"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("."))
  http.Handle("/", fs)

  log.Println("Listening:8080...")
  http.ListenAndServe(":8080", nil)
}