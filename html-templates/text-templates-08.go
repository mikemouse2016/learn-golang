/*
 Predefined template functions
 There are also a few predefined template functions that you can employ within your code. I shall illustrate the
 printf function which works similar to the function fmt.Sprintf.
*/

package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("test")
	t = template.Must(t.Parse("{{with $x := `hello`}}{{printf `%s %s` $x `Mary`}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)
}