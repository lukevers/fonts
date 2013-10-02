package main

import (
	"net/http"
	"net/url"
	"os"
)

// HandleCSS is an API func that serves CSS when requested to do
// so. It doesn't serve CSS to fonts that we do not have, it 404's
// them instead.
func HandleCSS(w http.ResponseWriter, r *http.Request) {
	v, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		l.Emergf("Could not parse query: %s", err)
	}
	
	// Check if font exists and send 404 if it does not.
	f := "googlefontdirectory/fonts/"+v.Get("family")+".ttf"
	if _, err := os.Stat(f); os.IsNotExist(err) {
		// Does not exist, send 404
		http.Error(w, "404 Not Found", http.StatusNotFound)
	} else {
		// File exists, serve css
		content := `@font-face {
  font-family: `+v.Get("family")+`;
  src: url(`+Conf.URL+`/fonts/`+v.Get("family")+`.woff) format('woff'), url(`+Conf.URL+`/fonts/`+v.Get("family")+`.ttf) format('ttf'), url(`+Conf.URL+`/fonts/`+v.Get("family")+`.otf) format('otf'), url(`+Conf.URL+`/fonts/`+v.Get("family")+`.eot) format('eot'), url(`+Conf.URL+`/fonts/`+v.Get("family")+`.svg) format('svg');
}`
		w.Header().Set("Content-Type", "text/css")
		w.Write([]byte(content))
	}
}