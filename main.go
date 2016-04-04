package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	short := vars["short"]

	req, err := http.NewRequest("GET", "http://goo.gl/"+short, nil)
	if err != nil {
		log.Fatal(err)
	}

	transport := http.Transport{}
	resp, err := transport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 301 {
		log.Fatal("Failed with status", resp.Status)
	}

	location := resp.Header.Get("Location")

	log.Println("Request " + short)

	http.Redirect(w, r, location, 301)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{short}", RedirectHandler)

	log.Println("Server started on :80")
	http.ListenAndServe(":80", r)
}
