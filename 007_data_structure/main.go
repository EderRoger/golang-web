package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.1.gohtml"))
}

func main() {
	//sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	sagesMap := map[string]string{
		"India":   "Gandhi",
		"America": "MLK",
		"Mediate": "Buddha",
		"Love":    "Jesus",
		"Prophet": "Muhammad",
	}

	err := tpl.Execute(os.Stdout, sagesMap)
	if err != nil {
		log.Fatalln(err)
	}

}
