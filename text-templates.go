// https://gobyexample.com/text-templates
//
// text/template allows us to create dynamic content, or show custom output to the user.
// there's also html/template, with the same API but extra security features for HTML.
//
// This is all rather familiar - because it's how Helm templates work.

package main

import (
	"os"
	"text/template"
)

func main() {

	// Create a new template and parse its body from a string.
	// Templates are a mixture of text and "actions" enclosed in {{...}}; these are used to dynamically insert content.
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// We can also use template.Must to panic if Parse returns an error.
	// This is especially useful for templates at global scope.
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// By "executing" the template, we generate its text with specific values for its actions.
	// The {{.}} action is replaced with the value passed as a parameter to Execute.
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// Define a helper function we use below.
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// If the data is a struct, we can use the {{.FieldName}} action to access its fields.
	// The fields should be exported to be accessible when a template is executing.
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// The same applies to maps; there is no restriction on the case of key names.
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// We can use if/else to provide conditional execution for templates.
	// A value is false if it's the default value of a type, like 0, an empty string, nil pointer etc
	// This example shows another feature of templates; using - in actions to trim whitespace.
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// range lets us loop through slices, arrays, maps or channels.
	// inside the range block {{.}} is set to the current item of the iteration.
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
