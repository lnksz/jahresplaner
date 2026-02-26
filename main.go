package main

import (
	"bytes"
	_ "embed"
	"log"
	"net/http"
	"time"
)

//go:embed index.html
var indexHTML []byte

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && r.URL.Path != "/index.html" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		http.ServeContent(w, r, "index.html", time.Time{}, bytes.NewReader(indexHTML))
	})

	log.Printf("Serving on http://localhost:2026")
	log.Fatal(http.ListenAndServe(":2026", mux))
}
