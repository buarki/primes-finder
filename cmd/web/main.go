package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/buarki/gowasm/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("env var PORT must be informed")
	}
	fmt.Println("Now Listening on ", port)
	publicDir, err := filepath.Abs("./public/")
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.Dir(publicDir))
	http.Handle("/public/", http.StripPrefix("/public", fs))
	http.HandleFunc("/", api.Handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
