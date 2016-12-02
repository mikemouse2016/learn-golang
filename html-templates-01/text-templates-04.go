/*
 pipelines
 Unix users will already be aware of ‘piping’ data. Many commands produce a textual output - a stream of text.
 If you type the command ls at the prompt, you will get a list of files in the directory.
 This can translated as get a list of all files in the current directory, and then pipe it to the default output
 which in this case is the command line screen. At the unix command line the pipe ‘|’ symbol which is
 the vertical line, allows you to ‘pipe’ the text stream to another command. For example,

 ls | grep "a"
 will get all the files in the local directory, and the pipe it to the grep command which filters only
 those files that contain the letter ‘a’ in it. Of course you could further pipe this text stream to another command.
 ls | grep "a" | grep "o", will list only those files with both an ‘a’ and an ‘o’ in it.

 In Go, every such text stream is called a pipeline and it can also be piped to other commands.
 In the example below we are simply printing two pipelines with constant strings.
 Note that these constant strings which are within {{ }} is different from static text outside of braces
 - the static text is copied as is without changes always.
 The pipeline data on the other hand could be manipulated, even though in this particular example we are not
 doing anything.
*/


package main

import (
	"text/template"
	"os"
)

func main() {
	t := template.New("template test")
	t = template.Must(t.Parse("This is just static text. \n{{\"This is pipeline data - because it is" +
		" evaluated within the double braces.\"}} {{`So is this, but within reverse quotes.`}}\n"));
	t.Execute(os.Stdout, nil)
}