package main

import (
	"io"
	"log"
	"os"
	"strings"
)

var name string

func main() {

	name = "Hello World"
	html := `<html><h2>` + name + `</h2></html>`

	file, err := os.Create("index.html")

	if err != nil {
		log.Fatal("error creating file", err)
	}

	defer file.Close()

	io.Copy(file, strings.NewReader(html))

}
