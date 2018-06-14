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

type car struct {
	Model string
	Door  int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tplStructSlice.gohtml"))
}

func main() {
	//sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love",
	}

	buddha := sage{
		Name:  "Buddha",
		Motto: "Peace",
	}

	corolla := car{
		Model: "Corolla",
		Door:  4,
	}

	civic := car{
		Model: "Civic",
		Door:  5,
	}

	sages := []sage{jesus, buddha}
	cars := []car{corolla, civic}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	dataRefact := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}
	// sagesMap := map[string]string{
	// 	"India":   "Gandhi",
	// 	"America": "MLK",
	// 	"Mediate": "Buddha",
	// 	"Love":    "Jesus",
	// 	"Prophet": "Muhammad",
	// }

	err := tpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, dataRefact)
	if err != nil {
		log.Fatalln(err)
	}

}
