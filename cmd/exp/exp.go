package main

import (
	"fmt"
	"html/template"
	"os"
)

type TemplateData struct {
	Nome string
	Age  int
}

func main() {

	t, err := template.ParseFiles("layout1.html", "footer.html", "header.html")

	fmt.Println(t.DefinedTemplates())

	fmt.Println(t.Name())

	if err != nil {
		panic(err)
	}

	// data := TemplateData{Nome: "Robson", Age: 17}

	err = t.Execute(os.Stdout, nil) // executa o template corrente
	if err != nil {
		panic(err)
	}
}
