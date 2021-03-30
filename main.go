package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var buildPath string

func main() {
	buildPath = os.Args[1]
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(serveSpa)))
}

func serveSpa(w http.ResponseWriter, r *http.Request) {
	serveFileOrIndex(w, r, getFilepath(r.URL.Path))
}

func serveFileOrIndex(w http.ResponseWriter, r *http.Request, requestedFilepath string) {
	if !fileExists(requestedFilepath) {
		http.ServeFile(w, r, getFilepath("index.html"))
		return
	}

	http.ServeFile(w, r, requestedFilepath)
}

func fileExists(requestedFilepath string) bool {
	_, err := os.Stat(requestedFilepath)
	if err != nil {
		return false
	}
	return true
}

func getFilepath(path string) string {
	return filepath.Join(buildPath, path)
}
