package static

import (
	"html/template"
	"strings"
)

func js(a asset) template.JS {
	return template.JS(a.Content)
}

func css(a asset) template.CSS {
	return template.CSS(a.Content)
}

func tmpl(a asset) *template.Template {

	funcMap := template.FuncMap{
		"toLower": strings.ToLower,
	}

	return template.Must(template.New(a.Name).Funcs(funcMap).Parse(a.Content))
}
