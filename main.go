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

func domainsearch(b string) string {
	domains := [3] string{"www.google.com", "facebook.com", "yahoo.com"}
	fmt.Println(b)
	for _, name := range domains  {
		fmt.Println(name)
		if name == b {
			return("Denied")
		} else {
			return("Ok")
		}
	} 
    return("Ok")
}

func proxyserver(w http.ResponseWriter, r *http.Request) {
	domain_name := r.Host
//	w.Write([]byte(domain_name))
	status := domainsearch(domain_name)

	fmt.Fprintf(w, "%v\n", status)
}

func main() {
	mux := http.NewServeMux()
	rik := http.HandlerFunc(proxyserver)
	mux.Handle("/", rik)
	http.ListenAndServe(":80", mux)
}
