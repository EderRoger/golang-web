package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
var tpl2 *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	tpl2 = template.Must(template.ParseFiles("tpl2.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 32)

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl2.ExecuteTemplate(os.Stdout, "tpl2.gohtml", `This is a message to template! Search knowledge`)

	if err != nil {
		log.Fatalln(err)
	}

}
