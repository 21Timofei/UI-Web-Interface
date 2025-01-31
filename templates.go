package templates

import (
	"embed"
	"html/template"
)

var TemplateFS embed.FS

func LoadTemplates() (*template.Template, error) {
	return template.ParseFS(TemplateFS, "templates/template.html")
}
