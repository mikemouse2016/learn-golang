/*
http://golangtutorials.blogspot.ro/2011/06/go-templates.html

 Typical usage of templates is within html code that is generated from the server side. We would open a template
 file which has already been defined, then merge that with some data we have using template.Execute which writes
 out the merged data into the io.Writer which is the first parameter. In the case of web functions,
 the io.Writer instance is passed into the handler as http.ResponseWriter.

*/

package main

import (
	"net/http"
	"html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {

	//create a new template
	t := template.New("some template")

	//open and parse a template text file
	t, _ = t.ParseFiles("tmpl/welcome.html", nil)

	//a method we have separately defined to get the value for a type
	user := GetCurrentlyLoggedInUser()

	//substitute fields in the template 't', with values from 'user'
	//and write it out to 'w' which implements io.Writer
	t.Execute(w, user)
}

func main() {

}

/*
 We should be seeing other examples of actual html code which utilizes this functionality.
 But for the purposes of learning, to write out all that html is unnecessary clutter. So for learning purposes, we shall use simpler code where I can illustrate the template concepts more clearly.

 * Instead of template.ParseFiles to which we have to pass one or more file paths, I shall use template.Parse
 for which I can give the text string directly in the code where it would be easier for you to see.
 * Instead of writing it as a web service, we shall write code we can execute from the command line
 * We shall use the predefined variable os.Stdout which refers to the standard output to print out the merged data
 - os.Stdout implements io.Writer
*/