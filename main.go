package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var buildPath string

func main() {
	buildPath = os.Getenv("REACT_BUILD_PATH")
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(serveSpa)))
}

func serveSpa(w http.ResponseWriter, r *http.Request) {
	serveFileOrIndex(w, r, getFilepath(r.URL.Path))
}

func getFilepath(path string) string {
	return filepath.Join(buildPath, path)
}

func serveFileOrIndex(w http.ResponseWriter, r *http.Request, filepath string) {
	if _, err := os.Stat(filepath); err != nil {
		http.ServeFile(w, r, getFilepath("index.html"))
		return
	}

	http.ServeFile(w, r, filepath)
}
