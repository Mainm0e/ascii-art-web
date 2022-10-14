package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on :8080...")
	http.ListenAndServe(":8080",nil)
}

func handlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Welcome to Ascii-Art-Web<h1>")
}