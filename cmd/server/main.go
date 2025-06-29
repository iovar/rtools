package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const DIR = "./dist"

func getFullPath(path string) string {
	return fmt.Sprintf("%s%c%s", DIR, os.PathSeparator, path)
}

func serveStaticAssets(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		return
	}

	path := req.URL.Path

	if path == "/" {
		path = "/index.html"
	}

	if strings.HasSuffix(path, ".css") {
		w.Header().Add("content-type", "text/css; charset=utf-8")
	}

	fullpath := getFullPath(path)

	file, err := os.ReadFile(fullpath)

	if err != nil {
		w.WriteHeader(404)
	}

	w.Write(file)
}

func main() {
	fmt.Println("Starting server on port 8080")

	http.HandleFunc("/", serveStaticAssets)

	http.ListenAndServe(":8080", nil)
}
