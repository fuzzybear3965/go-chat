package static

import "html/template"

func js(a asset) string {
	return a.Content
}

func css(a asset) string {
	return a.Content
}

func tmpl(a asset) *template.Template {
	return template.Must(template.New(a.Name).Parse(a.Content))
}
