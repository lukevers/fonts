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
	
/*	content := `@font-face {
  font-family: `+file+`;
  src: url(`+Conf.URL+`/fonts/`+file+`.ttf) format('ttf');
}`*/
	w.Header().Set("Content-Type", "text/css")
	w.Write([]byte(file))
}