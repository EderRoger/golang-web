package main

import (
	"log"
	"os"
	"text/template"
)

type breakfast struct {
	Name string
}

type lunch struct {
	Name string
}

type dinner struct {
	Items []string
}

type menu struct {
	Breakfast breakfast
	Lunch     lunch
	Dinner    dinner
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	restaurantMenu := menu{
		breakfast{
			Name: "Fruits",
		},
		lunch{
			Name: "Pasta with chrimps",
		},
		dinner{
			Items: []string{"Raw Meat", "Beens"},
		},
	}

	err := tpl.Execute(os.Stdout, restaurantMenu)
	if err != nil {
		log.Fatalln(err)
	}
}

// 1. Create a data structure to pass to a template which
// * contains information about restaurant's menu including Breakfast, Lunch, and Dinner items
