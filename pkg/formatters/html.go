package formatters

import (
	"io"
	"text/template"
)

func NewHtml() Formatter {
	return &htmlFormatter{}
}

type htmlFormatter struct {
}

func (f *htmlFormatter) ListPages(pageNames []string, writer io.Writer) {
	tmpl, _ := template.New("listing-template").Parse(listingTemplate)
	t := template.Must(tmpl, nil)
	t.Execute(writer, pageNames)
}

const listingTemplate = `
<html>
	<body>
		{{range $val := .}}
			<a href="https://en.wikipedia.org/wiki/{{$val}}">{{$val}}</a><br>
		{{end}}
	</body>
</html>
`
