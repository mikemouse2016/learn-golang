/*
 template Must function - to check validity of template text
 The static Must function checks for the validity of the template content, i.e. things like whether the braces
 are matches, whether comments are closed, and whether variables are properly formed, etc. In the example below,
 we have two valid template texts and they parse without causing a panic. The third one, however,
 has an unmatched brace and will panic.
*/


package main

import (
	"text/template"
	"fmt"
)

func main() {
	tOk := template.New("first")

	//a valid template, so no panic with Must
	template.Must(tOk.Parse(" some static text /* and a comment */"))
	fmt.Println("The first one parsed OK.")

	template.Must(template.New("second").Parse("some static text {{ .Name }}"))
	fmt.Println("The second one parsed OK.")

	fmt.Println("The next one ought to fail.")
	tErr := template.New("check parse error with Must")

	// due to unmatched brace, there should be a panic here
	template.Must(tErr.Parse(" some static text {{ .Name }"))
}