package api

import (
	"net/http"
)

const (
	wasmPath = "./public/main.wasm"
)

func Handler1(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, wasmPath)
}
