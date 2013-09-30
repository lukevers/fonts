package main

import (
	"net/http"
	"net/url"
)

func HandleCSS(w http.ResponseWriter, r *http.Request) {
	v, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		l.Emergf("Could not parse query: %s", err)
	}
	file := v.Get("name")
	// TODO
	// - [ ] Check if name exists in database
	// - [ ] Serve css if so
}