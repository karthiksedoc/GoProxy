package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)


func httpclient(b string) ([]byte, error) {
	response, err := http.Get(b)	
	if err != nil {
		//Need to handle error
		var my_response []byte
		return my_response, err
	}
	defer response.Body.Close()
	my_response, _ := ioutil.ReadAll(response.Body)
	return my_response, nil
}

func domainsearch(b string) ([]byte, string) {
	domains := [4] string{"www.google.com", "google.com", "facebook.com", "yahoo.com"}
//	fmt.Println(b)
	var web_response []byte
	for _, name := range domains  {
		fmt.Println(name)
		if name == b {
	//	   var web_response []byte
		   dom_status := "Denied"
		   return web_response, dom_status
		} 
	}
	webresponse, err := httpclient(b)
	if err != nil {
		return web_response, "Access Error"
	}
    return webresponse, ""
}

func proxyserver(w http.ResponseWriter, r *http.Request) {
	domain_name := r.Host
//	w.Write([]byte(domain_name))
	status, err := domainsearch(domain_name)

	if (err == "Denied") {
	   http.Error(w, "Blocked", 403)
	   return
	}

	fmt.Fprintf(w, "%v\n", status)
}

func main() {
	mux := http.NewServeMux()
	rik := http.HandlerFunc(proxyserver)
	mux.Handle("/", rik)
	http.ListenAndServe(":80", mux)
}
