package main

import (
	"html/template"
	"os"
)

type Data struct {
	Food        []string
	AnotherData []string
}

func main() {
	tmpl := `
	<ul>
		{{range $index, $element := .Food}}
			<li>Index: {{$index}}, Food: {{$element}}, AnotherData: {{index $.AnotherData $index}}</li>
		{{end}}
	</ul>
	`

	data := Data{
		Food:        []string{"rice", "burger", "pizza"},
		AnotherData: []string{"data1", "data2", "data3"},
	}

	t, err := template.New("example").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
