package main

import (
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main(){
  Buddha :=sage{
    Name: "Buddha",
    Motto: "The belief of no beliefs",
  }
  Gandhi :=sage{
    Name: "Gandhi",
    Motto: "The belief of no beliefs",
  }
  Mlk :=sage{
    Name: "Mlk",
    Motto: "The belief of no beliefs",
  }
  Jesus :=sage{
    Name: "Jesus",
    Motto: "The belief of no beliefs",
  }Muhammad :=sage{
    Name: "Muhammad",
    Motto: "The belief of no beliefs",
  }Buddha :=sage{
    Name: "Buddha",
    Motto: "The belief of no beliefs",
  }
  
}
