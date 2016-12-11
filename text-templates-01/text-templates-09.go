/*
A template Set WAS a specific data type that allowed you to group together related templates in one group.
Though it does not exist now as an individual data type, it is subsumed within the Template data structure.
So if you now parsed a set of files that contains text within it separated by {{define}}...{{end}} sections,
each of them become a template.

For example if a web page had a header, body, and footer, these could be defined within one text file.
This could then could be read into your program with one call which will parse all the template definitions.
For example, assume the text as shown below which illustrates template definitions within a single file.

Incomplete sample template set file
{{define "header"}}
<html>
<head></head>
{{end}}

{{define "body"}}
<body>
</body>
{{end}}

{{define "footer"}}
</html>
{{end}}


Key points to note:
* each template is enclosed within {{define}} and {{end}}
* each template is given a unique name - repeating the same name will cause a panic
* text outside of a {{define}} block is not allowed - it will cause a panic
* white space within the the {{define}} block is taken as is - in the above example there is a new line character
before the first html tag and also after the closing head tag.

When this particular file is parsed in as a set, it creates a map of template names to parsed templates.
So the above could be represented for illustration as):
tmplVar["header"] = pointer to parsed template of text "<html> … </head>"
tmplVar["body"] = pointer to parsed template of text "<body> … </body>"
tmplVar["footer"] = pointer to parsed template of text "</html>"


Templates within a set know about each other. If the same template needs to be used by different sets,
it needs to be parsed separately. If the "footer" above can be repeated, keep it as a separate file and have
each set parse it separately.


We shall now do a simple example. In this example, we shall learn how to
* define a template - using {{define}}
* include one template within another - using {{template "template name"}}
* parse a bunch of files and create a set using template.ParseFiles
* execute/merge contents of the templates using template.ExecuteTemplate
* and finally, how to add external templates to an existing set using template.Set.Add
(I don't think there is a way to do this anymore, but if somebody knows, please let me know too.)

Full template file - t1.tmpl
{{define "t_ab"}}a b{{template "t_cd"}}e f {{end}}

The file above will be parsed in as a template named "t_ab". It has, within it, "a b /missing/ e f", but is missing
a couple of letters in the alphabet. For that it intends to include another template called
"t_cd" (which should be in the same set).

Full template file - t2.tmpl
{{define "t_cd"}} c d {{end}}

The file above will be parsed in as a template called "t_cd".
*/

package main

import (
	"text/template"
	"os"
	"fmt"
)

func main() {
	fmt.Println("Load a set of templates with {{define}} clauses and execute:")

	//create a set of templates from many files.
	s1, _ := template.ParseFiles("t1.tmpl", "t2.tmpl")
	//Note that t1.tmpl is the file with contents "{{define "t_ab"}}a b{{template "t_cd"}}e f {{end}}"
	//Note that t2.tmpl is the file with contents "{{define "t_cd"}} c d {{end}}"

	//just printing of c d
	s1.ExecuteTemplate(os.Stdout, "t_cd", nil)
	fmt.Println()

	//execute t_ab which will include t_cd
	s1.ExecuteTemplate(os.Stdout, "t_ab", nil)
	fmt.Println()

	//since templates in this data structure are named, there is no default template and so it prints nothing
	s1.Execute(os.Stdout, nil)
}