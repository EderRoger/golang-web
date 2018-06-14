package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tplStruct.gohtml"))
}

func main() {
	//sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love",
	}

	// sagesMap := map[string]string{
	// 	"India":   "Gandhi",
	// 	"America": "MLK",
	// 	"Mediate": "Buddha",
	// 	"Love":    "Jesus",
	// 	"Prophet": "Muhammad",
	// }

	err := tpl.Execute(os.Stdout, jesus)
	if err != nil {
		log.Fatalln(err)
	}

}
