package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("ingredients/static/*.html"))
}

func main() {
	static := http.FileServer(http.Dir("./ingredients/static"))
	http.Handle("/", static)
	http.HandleFunc("/input_Ascii-Art", processor)
	http.ListenAndServe(":8000", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	texts := r.FormValue("inputText")
	d := struct {
		Dick  string
		Hello string
	}{
		Dick:  texts,
		Hello: " _    _          _   _\n| |  | |        | | | |\n| |__| |   ___  | | | |   ___\n|  __  |  / _ \\ | | | |  / _ \\\n| |  | | |  __/ | | | | | (_) |",
	}
	tpl.ExecuteTemplate(w, "index.html", d)

}
