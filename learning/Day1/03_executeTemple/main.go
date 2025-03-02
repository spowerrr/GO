package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	file, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Error creating file:", err)
	}
	defer nf.Close() //create ,os, close

	err = file.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Execution successfully done.")
	}
}
