/*
 The dot - .
 The dot (.) is used in Go templates to refer to the current pipeline. This is similar to a cursor used with
 database access to indicate the current row that is being used among all the rows returned by a query.
 Or if you are used to Java and C++, you could think it is something like the this operator - well, not really,
 but kinda.

 Along with certain constructs the value of the ‘dot’ gets automatically set to the current value in the pipeline.
 You can therefore refer to its value as {{.}}.

template with-end
The with statement sets the value of dot with the value of the pipeline. If the pipeline is empty,
then whatever is between the with-end block is skipped.
*/

package main

import (
	"os"
	"text/template"
)

func main() {
	t, _ := template.New("test").Parse("{{with `hello`}}{{.}}{{end}}!\n")
	t.Execute(os.Stdout, nil)

	//when nested, the dot takes the value according to closest scope.
	t1, _ := template.New("test").Parse("{{with `hello`}}{{.}} {{with `Mary`}}{{.}}{{end}}{{end}}!\n")
	t1.Execute(os.Stdout, nil)
}