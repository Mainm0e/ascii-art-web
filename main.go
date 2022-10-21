package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/mainm0e/asciiweb/rary"
)

var temp *template.Template

type Data struct {
	Output    string
	ErrorNum  int
	ErrorText string
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/ascii-art", AsciiArtHandle)
	fmt.Println("Start sever at port 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func AsciiArtHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	if r.URL.Path != "/ascii-art" {
		fmt.Println("hi")
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	//err 404 handle

	d := Data{}
	temp = template.Must(template.ParseGlob("static/*.html"))

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if r.Method == "GET" {
		temp.ExecuteTemplate(w, "index.html", d)
	} else if r.Method == "POST" {
		Ascii := r.FormValue("ascii-art")
		Font := r.FormValue("Font")
		d.Output = rary.Output(Ascii, Font)
	}
	temp.ExecuteTemplate(w, "resulindex.html", d)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
