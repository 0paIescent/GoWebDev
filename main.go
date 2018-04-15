package main

import (
	"html/template"
	"os"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type PageData struct {
	Title  string
	Header string
}

func renderIndex(pd PageData) {
	templates.ExecuteTemplate(os.Stdout, "index.gohtml", pd)
}

func main() {
	index := PageData{"Index - GoWebDev", "Welcome to Index!"}
	renderIndex(index)
}
