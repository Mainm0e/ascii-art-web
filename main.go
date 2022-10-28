package main

import (
	"fmt"
	"net/http"

	"github.com/mainm0e/asciiartwebstylize/docs/serverhandler"
)

func main() {
	Port := ":8000"

	http.HandleFunc("/", serverhandler.HomeHandler)
	http.HandleFunc("/ascii-art", serverhandler.ServerHandler)
	fileServer := http.FileServer(http.Dir("./docs"))
	http.Handle("/docs/", http.StripPrefix("/docs/", fileServer))
	fmt.Printf("Start sever at port %v...\n", Port)
	http.ListenAndServe(Port, nil)
}
