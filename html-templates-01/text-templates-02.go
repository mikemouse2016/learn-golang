/*
 For the sake of completeness, let’s do an example where there is an error due to a missing field.
 In the below code, we have a field nonExportedAgeField, which, since it starts with a small letter, is not exported.
 Therefore when merging there is an error. You can check for the error on the return value of the Execute function.
*/


package main



import (
	"os"
	"text/template"
	"fmt"
)

type Person struct {
	Name                string

	//because it doesn't start with a capital letter
	nonExportedAgeField string
}

func main() {
	p := Person{Name: "Mary", nonExportedAgeField: "31"}

	t := template.New("nonexported template demo")
	t, _ = t.Parse("hello {{.Name}}! Age is {{.nonExportedAgeField}}.")
	err := t.Execute(os.Stdout, p)
	if err != nil {
		fmt.Println("There was an error:", err)
	}
}