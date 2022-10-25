package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/mainm0e/asciiweb/docs/rary"
)

var temp *template.Template

type Data struct {
	Output    string
	ErrorNum  int
	ErrorText string
}

func main() {
	Port := ":8080"

	http.HandleFunc("/", ServerHandler)
	fileServer := http.FileServer(http.Dir("./docs"))
	http.Handle("/docs/", http.StripPrefix("/docs/", fileServer))
	fmt.Printf("Start sever at port %v...\n", Port)
	http.ListenAndServe(Port, nil)
}

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	d := Data{}
	temp = template.Must(template.ParseGlob("docs/static/*.html"))
	fmt.Println(r.URL.Path)
	if r.URL.Path != "/" {
		d.ErrorNum = 404
		d.ErrorText = "page Not Found"
		errorHandler(w, r, &d)
		return
	}
	Hometext, _ := rary.Output("Ascii Web", "standard.txt")
	d.Output = Hometext
	if r.Method == "GET" {
		temp.ExecuteTemplate(w, "index.html", d)
	} else if r.Method == "POST" {
		text := r.FormValue("input")
		font := r.FormValue("font")
		out, err := rary.Output(text, font)
		if err == false {
			d.ErrorNum = 500
			d.ErrorText = "Internal Server Error"
			errorHandler(w, r, &d)
			return
		}
		/* for making file
			FileName := "Text.pdf"
		output := []byte(NewString)
		MakeFile := os.WriteFile(FileName, output, 0644)
		Check(MakeFile)
		*/

		fmt.Println(out)
		d.Output = out
		temp.ExecuteTemplate(w, "index.html", d)
	} else {
		d.ErrorNum = 400
		d.ErrorText = "Bad Request"
		errorHandler(w, r, &d)
		return
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, d *Data) {
	w.WriteHeader(d.ErrorNum)
	temp.ExecuteTemplate(w, "err.html", d)
}
