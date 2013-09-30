package main

import (
	"net/http"
	"html/template"
)

var (
	t *template.Template
)

// Serve starts the HTTP server
func Serve(config *Config) {
	// Start the API if enabled
	if config.Access.EnableAPI {
		l.Info("Starting API")
		http.HandleFunc("/css", HandleCSS)
	}
	
	// Start the front-end if enabled
	if config.Access.EnableFrontEnd {
		l.Info("Starting front-end")
		http.HandleFunc("/", HandleRoot)
		http.HandleFunc("/assets", HandleAssets)
		// Compile templates
	//	t = template.Must(template.ParseGlob("templates/*.html"))
	}

	// Add the handler for the fonts
	http.HandleFunc("/fonts", HandleFonts)

	// Listen and serve
	http.ListenAndServe(config.Address, nil)
}

// HandleAssets is a static file server that serves everything in
// static directory
func HandleAssets(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

// HandleFonts is a static file server that serves all fonts
func HandleFonts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-font-ttf")
	http.ServeFile(w, r, "googlefontdirectory/"+r.URL.Path[1:])
}

// HandleRoot handles the "/" connections
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	
}