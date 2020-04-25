package main

import (
	"log"
	"net/http"
	"os"
)

const defaultPort = "8000"

func main() {
	// serve files
	http.Handle(
		"/",
		http.StripPrefix(
			"/",
			http.FileServer(http.Dir("static/html")),
		),
	)

	assetsDir := "static"
	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir(assetsDir)),
		),
	)

	wasmDir := "wasm"
	http.Handle(
		"/wasm/",
		http.StripPrefix(
			"/wasm/",
			http.FileServer(http.Dir(wasmDir)),
		),
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("Running with port %s", port)
	serve := http.ListenAndServe(":"+port, nil)
	if serve != nil {
		panic(serve)
	}
}
