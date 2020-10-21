package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)


func httpclient(b string) ([]byte, error) {
	fmt.Println(b)
	protocol := "http://"
	response, err := http.Get(protocol+b)
	fmt.Println(response)
	if err != nil {

		var my_response []byte
		return my_response, err
	}
	defer response.Body.Close()
	my_response, _ := ioutil.ReadAll(response.Body)

	fmt.Printf("httpclient func variable my_response = %T\n", my_response)
	return my_response, nil
}

func domainsearch(b string) ([]byte, string) {
	domains := [4] string{"www.google.com", "google.com", "facebook.com", "yahoo.com"}
	
	for _, name := range domains  {
		var web_response []byte
		if name == b {
		   dom_status := "Denied"
		   return web_response, dom_status
		} 
	}
	webresponse, err := httpclient(b)
	if err != nil {
		web_response := []byte(nil)
		return web_response, "Access Error"
	}
    return webresponse, ""
}

func proxyserver(w http.ResponseWriter, r *http.Request) {
	domain_name := r.Host
	status, err := domainsearch(domain_name)

	if (err == "Denied") {
	   http.Error(w, "Blocked", 403)
	   return
	}

	w.Write([]byte(status))
}

func main() {
	mux := http.NewServeMux()
	rik := http.HandlerFunc(proxyserver)
	mux.Handle("/", rik)
	http.ListenAndServe(":80", mux)
}
