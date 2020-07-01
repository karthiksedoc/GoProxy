package main

import (
	"net/http"
	"fmt"
)

//func proxyserver(domain string) http.Handler{
//	cust_handler := func(w http.ResponseWriter, r *http.Request) {
//		domain_name := r.URL
//		w.Write([]byte*(domain_name))
//	}
//	return http.HandlerFunc(cust_handler)
//}

func proxyserver(w http.ResponseWriter, r *http.Request) {
	domain_name := r.Host
//	w.Write([]byte(domain_name))
	fmt.Fprintf(w, "%v\n", domain_name)
}

func main() {
	mux := http.NewServeMux()
	rik := http.HandlerFunc(proxyserver)
	mux.Handle("/", rik)
	http.ListenAndServe(":80", mux)
}
