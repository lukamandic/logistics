package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/" {
		path = "/index.html"
	}

	path = strings.TrimPrefix(path, "/")

	staticPath := filepath.Join("static", path)

	if _, err := os.Stat(staticPath); os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join("static", "index.html"))
		return
	}

	ext := filepath.Ext(path)
	switch ext {
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".html":
		w.Header().Set("Content-Type", "text/html")
	case ".json":
		w.Header().Set("Content-Type", "application/json")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case ".svg":
		w.Header().Set("Content-Type", "image/svg+xml")
	case ".ico":
		w.Header().Set("Content-Type", "image/x-icon")
	}

	http.ServeFile(w, r, staticPath)
}