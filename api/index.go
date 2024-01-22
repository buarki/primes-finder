package api

import (
	"net/http"
)

const (
	htmlPath = "./public/index.html"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, htmlPath)
}
