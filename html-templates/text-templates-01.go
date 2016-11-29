/*
 Field substitution - {{.FieldName}}
 To include the content of a field within a template, enclose it within curly braces and add a dot at the beginning.
 E.g. if Name is a field within a struct and its value needs to be substituted while merging, then include the text
 {{.Name}} in the template. Do remember that the field name has to be present and it should also be exported
 (i.e. it should begin with a capital letter in the type definition), or there could be errors.
*/


package main

import (
	"os"
	"text/template"
)

type Person struct {
	//exported field since it begins with a capital letter
	Name string
}

func main() {

	//create a new template with some name
	t := template.New("hello template")

	//parse some content and generate a template, which is an internal representation
	t, _ = t.Parse("hello {{.Name}}!")

	//define an instance with required field
	p := Person{Name:"Mary"}

	//merge template ‘t’ with content of ‘p’
	t.Execute(os.Stdout, p)
}