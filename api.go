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

	content := `@font-face {
  font-family: `+v.Get("family")+`;
  src: url(`+Conf.URL+`/fonts/`+v.Get("family")+`.woff) format('woff'), url(`+Conf.URL+`/fonts/`+v.Get("family")+`.ttf) format('ttf'), url(`+Conf.URL+`/fonts/`+v.Get("family")+`.otf) format('otf'), url(`+Conf.URL+`/fonts/`+v.Get("family")+`.eot) format('eot'), url(`+Conf.URL+`/fonts/`+v.Get("family")+`.svg) format('svg');
}`
	w.Header().Set("Content-Type", "text/css")
	w.Write([]byte(content))
}