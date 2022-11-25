package template

import (
	"strings"
	"text/template"
)

func TemplateTest() {

	// t1 := template.New("t1")
	// t1, err := t1.Parse("Value is {{.}}\n")
	// if err != nil {
	// 	panic(err)
	// }

	// t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// t1.Execute(os.Stdout, "some text")
	// t1.Execute(os.Stdout, 5)
	// t1.Execute(os.Stdout, []string{
	// 	"Go",
	// 	"Rust",
	// 	"C++",
	// 	"C#",
	// })

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// t2 := Create("t2", "Name: {{.Name}}\n")

	// t2.Execute(os.Stdout, struct {
	// 	Name string
	// }{"Jane Doe"})

	// t2.Execute(os.Stdout, map[string]string{
	// 	"Name": "Mickey Mouse",
	// })

	t3 := Create("t3", `{{- if and (lt .cr_7d 0.7) (lt .rating_5d 4.5)  -}} Improve your rating and cr
{{ else if lt .rating_5d 4.5 -}} Improve your rating
{{ else -}} Improve your cr {{- end -}}`)
	b := new(strings.Builder)
	t3.Execute(b, map[string]float64{
		"cr_7d":     0.5,
		"rating_5d": 4.12,
	})
	t3.Execute(b, map[string]float64{
		"cr_7d":     0.9,
		"rating_5d": 4.12,
	})
	t3.Execute(b, map[string]float64{
		"cr_7d":     0.5,
		"rating_5d": 5.12,
	})

	println(b.String())

	// t4 := Create("t4",
	// 	"Range: {{range .}}{{.}} {{end}}\n")
	// t4.Execute(os.Stdout,
	// 	[]string{
	// 		"Go",
	// 		"Rust",
	// 		"C++",
	// 		"C#",
	// 	})
}
