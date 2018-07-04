package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotel := hotel{
		"Atrons",
		"Lake street",
		"California",
		"4938209",
		"CA",
	}

	err := tpl.Execute(os.Stdout, hotel)
	if err != nil {
		log.Fatalln(err)
	}

}

// 1. Create a data structure to pass to a template which
// * contains information about California hotels including Name, Address, City, Zip, Region
// * region can be: Southern, Central, Northern
// * can hold an unlimited number of hotels
