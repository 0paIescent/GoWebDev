package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	html, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = html.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
