package main

import (
  "net/http"
  "log"
)

func proxyHandler(w http.ResponseWriter, r *http.Request) {
  // function to match URL's from the DB
  // Write response to w based on whether URL matches or not
}

func main() {
  mux := http.NewServeMux()
  ph := http.HandlerFunc(proxyHandler)
  mux.Handle("/", ph)
  
  log.Println("Listening on port 80")
  http.ListenAndServe(":80", mux)
}
